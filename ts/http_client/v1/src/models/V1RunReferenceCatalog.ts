/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.0.6-rc0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface V1RunReferenceCatalog
 */
export interface V1RunReferenceCatalog {
    /**
     * 
     * @type {string}
     * @memberof V1RunReferenceCatalog
     */
    owner?: string;
    /**
     * 
     * @type {string}
     * @memberof V1RunReferenceCatalog
     */
    project?: string;
    /**
     * 
     * @type {string}
     * @memberof V1RunReferenceCatalog
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof V1RunReferenceCatalog
     */
    state?: string;
}

/**
 * Check if a given object implements the V1RunReferenceCatalog interface.
 */
export function instanceOfV1RunReferenceCatalog(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1RunReferenceCatalogFromJSON(json: any): V1RunReferenceCatalog {
    return V1RunReferenceCatalogFromJSONTyped(json, false);
}

export function V1RunReferenceCatalogFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1RunReferenceCatalog {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'owner': !exists(json, 'owner') ? undefined : json['owner'],
        'project': !exists(json, 'project') ? undefined : json['project'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'state': !exists(json, 'state') ? undefined : json['state'],
    };
}

export function V1RunReferenceCatalogToJSON(value?: V1RunReferenceCatalog | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'owner': value.owner,
        'project': value.project,
        'name': value.name,
        'state': value.state,
    };
}

