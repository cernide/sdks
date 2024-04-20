/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.1.8
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * - tpe: tpe algorithm
 *  - rand: random algorithm
 *  - anneal: anneal algorithm
 * @export
 */
export const V1HyperoptAlgorithms = {
    Tpe: 'tpe',
    Rand: 'rand',
    Anneal: 'anneal'
} as const;
export type V1HyperoptAlgorithms = typeof V1HyperoptAlgorithms[keyof typeof V1HyperoptAlgorithms];


export function V1HyperoptAlgorithmsFromJSON(json: any): V1HyperoptAlgorithms {
    return V1HyperoptAlgorithmsFromJSONTyped(json, false);
}

export function V1HyperoptAlgorithmsFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1HyperoptAlgorithms {
    return json as V1HyperoptAlgorithms;
}

export function V1HyperoptAlgorithmsToJSON(value?: V1HyperoptAlgorithms | null): any {
    return value as any;
}

