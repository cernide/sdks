/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.0.0-rc14
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * 
 * @export
 */
export const SparkDeployMode = {
    Cluster: 'cluster',
    Client: 'client',
    InClusterClient: 'in_cluster_client'
} as const;
export type SparkDeployMode = typeof SparkDeployMode[keyof typeof SparkDeployMode];


export function SparkDeployModeFromJSON(json: any): SparkDeployMode {
    return SparkDeployModeFromJSONTyped(json, false);
}

export function SparkDeployModeFromJSONTyped(json: any, ignoreDiscriminator: boolean): SparkDeployMode {
    return json as SparkDeployMode;
}

export function SparkDeployModeToJSON(value?: SparkDeployMode | null): any {
    return value as any;
}

