// React imports
import React, { ChangeEvent } from 'react';

// Library imports
import { CdsControlMessage } from '@cds/react/forms';
import { useFormContext } from 'react-hook-form';
import { CdsSelect } from '@cds/react/select';
import * as yup from 'yup';

export interface OsImage {
    // NOTE: we expect all images to have a name; however, the swagger contract generates objects where the name is defined as optional
    // so we define our interface to be compatible with an optional name attribute
    // NOTE: this name MUST be unique for all images in the array, because we use the name to find the corresponding image
    name?: string;
}

interface ImageProps {
    osImageLabel: string;
    field: string;
    onOsImageSelected: (osImage: OsImage | undefined, fieldName?: string) => void;
    images: OsImage[];
}
export function osImagesValidation() {
    return yup.string().nullable().required('Please select an OS image');
}

function OsImageSelect(props: ImageProps) {
    const { osImageLabel, field, images, onOsImageSelected } = props;
    const handleOsImageSelect = (event: ChangeEvent<HTMLSelectElement>) => {
        const selectedOsImage = props.images.find((osImage) => osImage.name === event.target.value);
        onOsImageSelected(selectedOsImage, field);
    };
    const {
        register,
        formState: { errors },
    } = useFormContext();

    const fieldError = errors[field];
    return (
        <div cds-layout="m:lg">
            <CdsSelect layout="vertical" controlWidth="shrink">
                <label>{osImageLabel}</label>
                <select {...register(field)} onChange={handleOsImageSelect}>
                    {images.map((image) => (
                        <option key={image.name}>{image.name}</option>
                    ))}
                </select>
            </CdsSelect>
            {fieldError && (
                <div>
                    &nbsp;<CdsControlMessage status="error">{fieldError.message}</CdsControlMessage>
                </div>
            )}
        </div>
    );
}

export default OsImageSelect;
