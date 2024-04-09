/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.1.6-rc0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * - roc: ROC curve
 *  - pr: Precision Recall curve
 *  - custom: Custom curve
 * @export
 */
export const V1EventCurveKind = {
    Roc: 'roc',
    Pr: 'pr',
    Custom: 'custom'
} as const;
export type V1EventCurveKind = typeof V1EventCurveKind[keyof typeof V1EventCurveKind];


export function V1EventCurveKindFromJSON(json: any): V1EventCurveKind {
    return V1EventCurveKindFromJSONTyped(json, false);
}

export function V1EventCurveKindFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1EventCurveKind {
    return json as V1EventCurveKind;
}

export function V1EventCurveKindToJSON(value?: V1EventCurveKind | null): any {
    return value as any;
}

