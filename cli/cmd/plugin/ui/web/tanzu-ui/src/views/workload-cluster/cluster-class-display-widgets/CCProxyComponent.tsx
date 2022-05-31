import { CCVariable, ClusterClassVariableType } from '../../../shared/models/ClusterClass';
import { populateDefaults } from './CCUtil';

export function createProxyComponentCCVar(defaults: any): CCVariable {
    return populateDefaults(defaults, ProxyComponentCCVar());
}

function ProxyComponentCCVar(): CCVariable {
    return {
        name: 'proxy',
        label: 'Proxy',
        info: 'Use this panel to configure the proxy information of your proxy server and what IPs should be routed to it.',
        // taxonomy: ClusterClassVariableType.GROUP_OPTIONAL,
        taxonomy: ClusterClassVariableType.GROUP,
        required: false,
        children: ProxyComponentChildren(),
    };
}

function ProxyComponentChildren(): CCVariable[] {
    return [
        {
            name: 'httpProxy',
            prompt: 'Use proxy for HTTP calls to these IPs (list):',
            taxonomy: ClusterClassVariableType.IP_LIST,
        },
        {
            name: 'httpsProxy',
            prompt: 'Use proxy for HTTPS calls to these IPs (list):',
            taxonomy: ClusterClassVariableType.IP_LIST,
        },
        {
            name: 'noProxy',
            prompt: 'Use NO proxy for calls to these IPs (list):',
            taxonomy: ClusterClassVariableType.IP_LIST,
        },
        {
            name: 'proxyCA',
            prompt: 'Certificate for proxy',
            taxonomy: ClusterClassVariableType.STRING_PARAGRAPH,
        },
    ];
}
