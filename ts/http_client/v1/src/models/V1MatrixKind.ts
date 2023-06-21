/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc21
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
export const V1MatrixKind = {
    Random: 'random',
    Grid: 'grid',
    Hyperband: 'hyperband',
    Bayes: 'bayes',
    Hyperopt: 'hyperopt',
    Iterative: 'iterative',
    Mapping: 'mapping'
} as const;
export type V1MatrixKind = typeof V1MatrixKind[keyof typeof V1MatrixKind];


export function V1MatrixKindFromJSON(json: any): V1MatrixKind {
    return V1MatrixKindFromJSONTyped(json, false);
}

export function V1MatrixKindFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1MatrixKind {
    return json as V1MatrixKind;
}

export function V1MatrixKindToJSON(value?: V1MatrixKind | null): any {
    return value as any;
}

