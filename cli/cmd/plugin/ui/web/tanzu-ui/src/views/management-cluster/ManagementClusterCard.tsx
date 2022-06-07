// React imports
import React, { useState } from 'react';

// Library imports
import { CdsButton } from '@cds/react/button';
import { CdsCard } from '@cds/react/card';
import { CdsDivider } from '@cds/react/divider';

interface ManagementClusterCardProps {
    name: string;
    context: string;
    description: string;
}

function ManagementClusterCard(props:ManagementClusterCardProps) {
    return (
        <div cds-layout="grid cols:12">
            <CdsCard>
                <div cds-layout="vertical gap:md" aria-labelledby="{id}">
                    <h2 id="{id}" cds-text="section">
                        Card Title
                    </h2>

                    <CdsDivider cds-card-remove-margin></CdsDivider>

                    <div cds-text="body light">Card Content</div>

                    <CdsDivider cds-card-remove-margin></CdsDivider>

                    <div cds-layout="horizontal gap:sm align:vertical-center">
                        <CdsButton action="flat-inline">View</CdsButton>
                    </div>
                </div>
            </CdsCard>
        </div>
    );
}

export default ManagementClusterCard;
