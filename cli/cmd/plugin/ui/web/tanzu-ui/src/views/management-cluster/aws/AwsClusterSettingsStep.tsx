// React imports
import React, { useContext, useEffect, useState } from 'react';

// Library imports
import { CdsButton } from '@cds/react/button';
import { ClarityIcons, blockIcon, blocksGroupIcon, clusterIcon } from '@cds/core/icon';
import { FormProvider, useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup/dist/yup';
import * as yup from 'yup';

// App imports
import { AwsStore } from './store/Aws.store.mc';
import { AWSVirtualMachine } from '../../../swagger-api';
import { AWS_FIELDS } from './aws-mc-basic/AwsManagementClusterBasic.constants';
import { ClusterName, clusterNameValidation } from '../../../shared/components/FormInputComponents/ClusterName/ClusterName';
import { FormAction } from '../../../shared/types/types';
import { getResource } from '../../../views/providers/aws/AwsResources.reducer';
import { INPUT_CHANGE } from '../../../state-management/actions/Form.actions';
import { NodeProfile, nodeProfileValidation } from '../../../shared/components/FormInputComponents/NodeProfile/NodeProfile';
import OsImageSelect from '../../../shared/components/FormInputComponents/OsImageSelect/OsImageSelect';
import { StepProps } from '../../../shared/components/wizard/Wizard';
import { STORE_SECTION_FORM } from '../../../state-management/reducers/Form.reducer';
import UseUpdateTabStatus from '../../../shared/components/wizard/UseUpdateTabStatus.hooks';
import { nodeProfiles } from './aws-mc-common/aws-orchestrator/AwsOrchestrator.service';

ClarityIcons.addIcons(blockIcon, blocksGroupIcon, clusterIcon);

type AWS_CLUSTER_SETTING_STEP_FIELDS = AWS_FIELDS.NODE_PROFILE | AWS_FIELDS.OS_IMAGE | AWS_FIELDS.CLUSTER_NAME;
interface AwsClusterSettingFormInputs {
    [AWS_FIELDS.CLUSTER_NAME]: string;
    [AWS_FIELDS.NODE_PROFILE]: string;
    [AWS_FIELDS.OS_IMAGE]: string;
}

function yupStringRequired(errorMessage: string) {
    return yup.string().nullable().required(errorMessage);
}

function createYupSchemaObject() {
    return {
        [AWS_FIELDS.OS_IMAGE]: yupStringRequired('Please select an OS image'),
        [AWS_FIELDS.NODE_PROFILE]: nodeProfileValidation(),
        [AWS_FIELDS.CLUSTER_NAME]: clusterNameValidation(),
    };
}

function AwsClusterSettingsStep(props: Partial<StepProps>) {
    const { updateTabStatus, currentStep, goToStep, submitForm } = props;
    const { awsState, awsDispatch } = useContext(AwsStore);
    const awsClusterSettingsFormSchema = yup.object(createYupSchemaObject()).required();
    const methods = useForm<AwsClusterSettingFormInputs>({
        resolver: yupResolver(awsClusterSettingsFormSchema),
        mode: 'all',
    });
    const {
        handleSubmit,
        formState: { errors },
        setValue,
    } = methods;

    // update tab status bar
    if (updateTabStatus) {
        UseUpdateTabStatus(errors, currentStep, updateTabStatus);
    }

    const goToNextStep = () => {
        if (goToStep && submitForm && currentStep) {
            goToStep(currentStep + 1);
            submitForm(currentStep);
        }
    };

    const onFieldChange = (field: AWS_CLUSTER_SETTING_STEP_FIELDS, data: any) => {
        awsDispatch({
            type: INPUT_CHANGE,
            field,
            payload: data,
        } as FormAction);
    };

    const osImages = (getResource(AWS_FIELDS.OS_IMAGE, awsState) || []) as AWSVirtualMachine[];
    let initialSelectedNodeProfileId = awsState[STORE_SECTION_FORM][AWS_FIELDS.NODE_PROFILE];

    if (awsState[STORE_SECTION_FORM][AWS_FIELDS.OS_IMAGE]) {
        setValue(AWS_FIELDS.OS_IMAGE, awsState[STORE_SECTION_FORM][AWS_FIELDS.OS_IMAGE].name);
    }

    if (!initialSelectedNodeProfileId) {
        initialSelectedNodeProfileId = nodeProfiles[0].id;
        setValue(AWS_FIELDS.NODE_PROFILE, initialSelectedNodeProfileId);
    }

    useEffect(() => {
        onFieldChange(AWS_FIELDS.NODE_PROFILE, initialSelectedNodeProfileId);
    }, []);

    const [selectedNodeProfileId, setSelectedNodeProfileId] = useState(initialSelectedNodeProfileId);

    const onInstanceTypeChange = (instanceType: string) => {
        setSelectedNodeProfileId(instanceType);
        onFieldChange(AWS_FIELDS.NODE_PROFILE, instanceType);
    };

    return (
        <FormProvider {...methods}>
            <div className="cluster-settings-container" cds-layout="p:lg">
                <h3 cds-layout="m-t:md m-b:xl" cds-text="title">
                    Management Cluster settings
                </h3>
                <div cds-layout="grid gap:m" key="section-holder">
                    <div cds-layout="col:4" key="cluster-name-section">
                        <ClusterName
                            field={AWS_FIELDS.CLUSTER_NAME}
                            clusterNameChange={(value) => {
                                onFieldChange(AWS_FIELDS.CLUSTER_NAME, value);
                            }}
                            placeholderClusterName={'my-aws-cluster'}
                            defaultClusterName={awsState[STORE_SECTION_FORM][AWS_FIELDS.CLUSTER_NAME]}
                        />
                    </div>
                    <div cds-layout="col:8" key="instance-type-section">
                        <NodeProfile
                            field={AWS_FIELDS.NODE_PROFILE}
                            nodeInstanceTypes={nodeProfiles}
                            nodeInstanceTypeChange={onInstanceTypeChange}
                            selectedInstanceId={selectedNodeProfileId}
                        />
                    </div>
                    <div cds-layout="col:12">
                        <OsImageSelect
                            osImageLabel="OS Image with Kubernetes"
                            images={osImages}
                            field={AWS_FIELDS.OS_IMAGE}
                            onOsImageSelected={(value) => {
                                onFieldChange(AWS_FIELDS.OS_IMAGE, value);
                            }}
                        />
                    </div>

                    <div cds-layout="grid col:12 p-t:lg">
                        <CdsButton onClick={handleSubmit(goToNextStep)}>NEXT</CdsButton>
                    </div>
                </div>
            </div>
        </FormProvider>
    );
}

export default AwsClusterSettingsStep;