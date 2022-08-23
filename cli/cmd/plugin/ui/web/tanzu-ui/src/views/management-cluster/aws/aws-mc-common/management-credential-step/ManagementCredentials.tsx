// React imports
import React, { ChangeEvent, MouseEvent, useContext, useEffect, useState } from 'react';

// Library imports
import { CdsButton } from '@cds/react/button';
import { CdsIcon } from '@cds/react/icon';
import { ClarityIcons, refreshIcon, connectIcon, infoCircleIcon } from '@cds/core/icon';
import { CdsControlMessage, CdsFormGroup } from '@cds/react/forms';
import { CdsRadioGroup, CdsRadio } from '@cds/react/radio';
import { FormProvider, SubmitHandler, useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';

// App import
import { AwsOrchestrator } from '../aws-orchestrator/AwsOrchestrator.service';
import { AwsService } from '../../../../../swagger-api';
import { AwsStore } from '../../store/Aws.store.mc';
import { AWSAccountParams } from '../../../../../swagger-api/models/AWSAccountParams';
import { AWSKeyPair } from '../../../../../swagger-api/models/AWSKeyPair';
import { AWS_FIELDS, CREDENTIAL_TYPE } from '../../aws-mc-basic/AwsManagementClusterBasic.constants';
import ConnectionNotification, { CONNECTION_STATUS } from '../../../../../shared/components/ConnectionNotification/ConnectionNotification';
import { FormAction } from '../../../../../shared/types/types';
import { INPUT_CHANGE } from '../../../../../state-management/actions/Form.actions';
import { managementCredentialFormSchema } from './management.credential.form.schema';
import ManagementCredentialOneTime from './ManagementCredentialOneTime';
import ManagementCredentialProfile from './ManagementCredentialProfile';
import SpinnerSelect from '../../../../../shared/components/Select/SpinnerSelect';
import { StepProps } from '../../../../../shared/components/wizard/Wizard';
import { STORE_SECTION_FORM } from '../../../../../state-management/reducers/Form.reducer';
import UseUpdateTabStatus from '../../../../../shared/components/wizard/UseUpdateTabStatus.hooks';
import './ManagementCredentials.scss';

ClarityIcons.addIcons(refreshIcon, connectIcon, infoCircleIcon);

export interface FormInputs {
    [AWS_FIELDS.PROFILE]: string;
    [AWS_FIELDS.REGION]: string;
    [AWS_FIELDS.SECRET_ACCESS_KEY]: string;
    [AWS_FIELDS.SESSION_TOKEN]: string;
    [AWS_FIELDS.ACCESS_KEY_ID]: string;
    [AWS_FIELDS.EC2_KEY_PAIR]: string;
}

function ManagementCredentials(props: Partial<StepProps>) {
    const { currentStep, goToStep, submitForm, updateTabStatus } = props;
    const { awsState, awsDispatch } = useContext(AwsStore);
    const [connectionStatus, setConnectionStatus] = useState<CONNECTION_STATUS>(CONNECTION_STATUS.DISCONNECTED);
    const [message, setMessage] = useState('');
    const [keyPairLoading, setKeyPairLoading] = useState(false);

    const methods = useForm<FormInputs>({
        resolver: yupResolver(managementCredentialFormSchema),
        mode: 'all',
        defaultValues: {
            [AWS_FIELDS.SECRET_ACCESS_KEY]: '',
            [AWS_FIELDS.ACCESS_KEY_ID]: '',
            [AWS_FIELDS.SESSION_TOKEN]: '',
            [AWS_FIELDS.PROFILE]: '',
            [AWS_FIELDS.REGION]: '',
            [AWS_FIELDS.EC2_KEY_PAIR]: '',
        },
    });
    const {
        register,
        setValue,
        getValues,
        handleSubmit,
        formState: { errors },
    } = methods;

    const [type, setType] = useState<CREDENTIAL_TYPE>(CREDENTIAL_TYPE.PROFILE);

    const [keypairs, setKeyPairs] = useState<AWSKeyPair[]>([]);

    const [errorObject, setErrorObject] = useState<{ [key: string]: any }>({});

    // update tab status bar
    if (updateTabStatus) {
        UseUpdateTabStatus(errors, currentStep, updateTabStatus);
    }

    useEffect(() => {
        if (connectionStatus === CONNECTION_STATUS.CONNECTED) {
            const initEC2KeyPairs = async () => {
                setKeyPairLoading(true);
                await AwsOrchestrator.initEC2KeyPairs(
                    {
                        awsState,
                        awsDispatch,
                        errorObject,
                        setErrorObject,
                    },
                    setKeyPairs
                );
                setKeyPairLoading(false);
            };
            initEC2KeyPairs();
        }
    }, [connectionStatus]);

    useEffect(() => {
        AwsOrchestrator.initNodeProfile({ awsState, awsDispatch, errorObject, setErrorObject });
    }, []);

    const selectCredentialType = (event: ChangeEvent<HTMLSelectElement>) => {
        setConnectionStatus(CONNECTION_STATUS.DISCONNECTED);
        setType(event.target.value as CREDENTIAL_TYPE);
        // We set the credentials type in the store to use later in the review config step
        awsDispatch({
            type: INPUT_CHANGE,
            field: AWS_FIELDS.CREDENTIAL_TYPE,
            payload: event.target.value,
        } as FormAction);
    };

    const onSubmit: SubmitHandler<FormInputs> = (data) => {
        const valueList = getValues();
        for (const i in valueList) {
            awsDispatch({
                type: INPUT_CHANGE,
                field: i,
                payload: valueList[i],
            } as FormAction);
        }

        if (connectionStatus === CONNECTION_STATUS.CONNECTED && Object.keys(errors).length === 0) {
            if (goToStep && currentStep && submitForm) {
                goToStep(currentStep + 1);
                submitForm(currentStep);
            }
        }
    };

    const handleConnect = async () => {
        let params: AWSAccountParams = {};
        if (type === CREDENTIAL_TYPE.PROFILE) {
            params = {
                profileName: awsState[STORE_SECTION_FORM].PROFILE,
                region: awsState[STORE_SECTION_FORM].REGION,
            };
        } else {
            params = {
                accessKeyID: awsState[STORE_SECTION_FORM].ACCESS_KEY_ID,
                region: awsState[STORE_SECTION_FORM].REGION,
                secretAccessKey: awsState[STORE_SECTION_FORM].SECRET_ACCESS_KEY,
                sessionToken: awsState[STORE_SECTION_FORM].SESSION_TOKEN,
            };
        }
        try {
            setConnectionStatus(CONNECTION_STATUS.CONNECTING);
            setMessage('Connecting to AWS');
            await AwsService.setAwsEndpoint(params);
            setConnectionStatus(CONNECTION_STATUS.CONNECTED);
            setMessage('Connected to AWS');
            AwsOrchestrator.initOsImages({ awsState, awsDispatch, errorObject, setErrorObject });
        } catch (err: any) {
            setConnectionStatus(CONNECTION_STATUS.ERROR);
            setMessage(`Unable to connect to AWS: ${err.body.message}`);
        }
    };

    const handleSelectKeyPair = (event: ChangeEvent<HTMLSelectElement>) => {
        awsDispatch({
            type: INPUT_CHANGE,
            field: AWS_FIELDS.EC2_KEY_PAIR,
            payload: event.target.value,
        } as FormAction);
    };

    const resetEc2KeyPair = () => {
        setConnectionStatus(CONNECTION_STATUS.DISCONNECTED);
        if (awsState[STORE_SECTION_FORM][AWS_FIELDS.EC2_KEY_PAIR] !== '') {
            setValue(AWS_FIELDS.EC2_KEY_PAIR, '', { shouldValidate: true });
            awsDispatch({
                type: INPUT_CHANGE,
                field: AWS_FIELDS.EC2_KEY_PAIR,
                payload: '',
            } as FormAction);
        }
    };

    const handleRefresh = async (event: MouseEvent<HTMLAnchorElement>) => {
        event.preventDefault();
        if (connectionStatus === CONNECTION_STATUS.CONNECTED) {
            const initEC2KeyPairs = async () => {
                setKeyPairLoading(true);
                await AwsOrchestrator.initEC2KeyPairs(
                    {
                        awsState,
                        awsDispatch,
                        errorObject,
                        setErrorObject,
                    },
                    setKeyPairs
                );
                setKeyPairLoading(false);
            };
            initEC2KeyPairs();
        }
    };

    function showErrorInfo() {
        if (connectionStatus === CONNECTION_STATUS.CONNECTED && JSON.stringify(errorObject) !== '{}') {
            return (
                <div>
                    <div className="error-text">Error Occurred</div>
                    <br />
                    {Object.keys(errorObject).map((errorField) => {
                        return (
                            <CdsControlMessage status="error" key={errorField}>
                                {errorObject[errorField]}
                            </CdsControlMessage>
                        );
                    })}
                    <br />
                </div>
            );
        }
        return;
    }

    const canContinue = (): boolean => {
        return (
            connectionStatus === CONNECTION_STATUS.CONNECTED &&
            getValues(AWS_FIELDS.EC2_KEY_PAIR) !== undefined &&
            JSON.stringify(errorObject) === '{}'
        );
    };

    return (
        <div className="wizard-content-container">
            <h2 cds-layout="m-t:md m-b:xl" cds-text="title">
                Amazon Web Services Credentials
            </h2>
            <CdsRadioGroup layout="vertical-inline" onChange={selectCredentialType}>
                <label cds-text="section medium" cds-layout="m-b:md">
                    Credential Type
                </label>
                <CdsRadio>
                    <label cds-layout="p-r:xxl">AWS credential profile</label>
                    <input type="radio" value={CREDENTIAL_TYPE.PROFILE} checked={type === CREDENTIAL_TYPE.PROFILE} readOnly />
                </CdsRadio>
                <CdsRadio>
                    <label>One-time credential</label>
                    <input type="radio" value={CREDENTIAL_TYPE.ONE_TIME} checked={type === CREDENTIAL_TYPE.ONE_TIME} readOnly />
                </CdsRadio>
            </CdsRadioGroup>
            <FormProvider {...methods}>
                {type === CREDENTIAL_TYPE.PROFILE && <ManagementCredentialProfile selectCallback={resetEc2KeyPair} />}
                {type === CREDENTIAL_TYPE.ONE_TIME && <ManagementCredentialOneTime selectCallback={resetEc2KeyPair} />}
            </FormProvider>
            <CdsFormGroup layout="vertical-inline" control-width="shrink">
                <div cds-layout="p-t:lg" className="aws-button-container">
                    <CdsButton
                        onClick={handleConnect}
                        disabled={connectionStatus === CONNECTION_STATUS.CONNECTED || !awsState[STORE_SECTION_FORM][AWS_FIELDS.REGION]}
                    >
                        <CdsIcon shape="connect" size="md"></CdsIcon>
                        CONNECT
                    </CdsButton>
                    <ConnectionNotification message={message} status={connectionStatus}></ConnectionNotification>
                </div>
                <div cds-layout="horizontal gap:lg align:vertical-center">
                    <SpinnerSelect
                        className="select-md-width"
                        disabled={connectionStatus !== CONNECTION_STATUS.CONNECTED}
                        label="EC2 key pair"
                        handleSelect={handleSelectKeyPair}
                        name={AWS_FIELDS.EC2_KEY_PAIR}
                        controlMessage="EC2 key pairs will be retrieved when connected to AWS."
                        isLoading={keyPairLoading}
                        register={register}
                        error={errors[AWS_FIELDS.EC2_KEY_PAIR]?.message}
                    >
                        <option></option>
                        {keypairs.map((keypair) => (
                            <option key={keypair.id} value={keypair.name}>
                                {keypair.name}
                            </option>
                        ))}
                    </SpinnerSelect>
                    <a
                        href="/"
                        className={
                            connectionStatus === CONNECTION_STATUS.CONNECTED && !keyPairLoading
                                ? 'btn-refresh icon-blue'
                                : 'btn-refresh disabled'
                        }
                        onClick={handleRefresh}
                        cds-text="secondary"
                    >
                        <CdsIcon shape="refresh" size="sm"></CdsIcon>{' '}
                        <span cds-layout="m-t:sm" className="vertical-mid">
                            REFRESH
                        </span>
                    </a>
                    <div cds-layout="col:6 align:horizontal-center">{showErrorInfo()}</div>
                </div>

                <CdsButton disabled={!canContinue()} onClick={handleSubmit(onSubmit)}>
                    NEXT
                </CdsButton>
            </CdsFormGroup>
        </div>
    );
}

export default ManagementCredentials;