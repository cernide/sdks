/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.1.0-rc5
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
export const MXJobMode = {
    MxTrain: 'MXTrain',
    MxTune: 'MXTune'
} as const;
export type MXJobMode = typeof MXJobMode[keyof typeof MXJobMode];


export function MXJobModeFromJSON(json: any): MXJobMode {
    return MXJobModeFromJSONTyped(json, false);
}

export function MXJobModeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MXJobMode {
    return json as MXJobMode;
}

export function MXJobModeToJSON(value?: MXJobMode | null): any {
    return value as any;
}

