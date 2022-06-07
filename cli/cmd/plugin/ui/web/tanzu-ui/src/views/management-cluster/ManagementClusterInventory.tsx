// React imports
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

// Library imports
import { CdsButton } from '@cds/react/button';
import { CdsIcon } from '@cds/react/icon';

// App imports
import { ManagementCluster } from '../../swagger-api/models/ManagementCluster';
import ManagementClusterCard from './ManagementClusterCard';
import { NavRoutes } from '../../shared/constants/NavRoutes.constants';
import './ManagementClusterInventory.scss';

function ManagementClusterInventory() {
    const [managementClusters, setManagementClusters] = useState<ManagementCluster[]>([
        {
            name: 'mycluster',
            provider: 'aws',
            endpoint: 'some endpoint',
            description: 'some description',
        },
    ]);

    const navigate = useNavigate();

    return (
        <div className="management-cluster-landing-container" cds-layout="vertical gap:md col@sm:12 grid">
            <div cds-layout="vertical col:8 gap:lg">
                <div cds-text="title">
                    <CdsIcon cds-layout="m-r:sm" shape="cluster" size="xl" className="icon-blue"></CdsIcon>
                    Management Clusters
                </div>
                {managementClusters.length === 0 && (
                    <div cds-layout="grid horizontal cols:8 p:md" className="section-raised mgmt-cluster-no-cluster-container">
                        <div cds-layout="grid horizontal cols:12 gap:lg gap@md:lg">
                            <div cds-text="body">
                                Create a Management Cluster on your preferred cloud provider through a guided series of steps.
                                <br />
                                <br />
                                This cluster will manage new workload clusters you create for your workloads.
                            </div>
                        </div>
                    </div>
                )}
                <div>
                    <CdsButton size="sm" onClick={() => navigate(NavRoutes.MANAGEMENT_CLUSTER_SELECT_PROVIDER)}>
                        <CdsIcon shape="cluster"></CdsIcon>create a new management cluster
                    </CdsButton>
                </div>
                {managementClusters.length > 0 &&
                    managementClusters.map((cluster) => {
                        return <ManagementClusterCard ></ManagementClusterCard>;
                    })}
            </div>
            <div cds-layout="col:4" className="mgmt-cluster-admins-img"></div>
        </div>
    );
}

export default ManagementClusterInventory;
