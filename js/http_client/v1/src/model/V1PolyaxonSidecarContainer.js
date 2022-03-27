// Copyright 2018-2021 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.17.1
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The V1PolyaxonSidecarContainer model module.
 * @module model/V1PolyaxonSidecarContainer
 * @version 1.17.1
 */
class V1PolyaxonSidecarContainer {
    /**
     * Constructs a new <code>V1PolyaxonSidecarContainer</code>.
     * @alias module:model/V1PolyaxonSidecarContainer
     */
    constructor() { 
        
        V1PolyaxonSidecarContainer.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1PolyaxonSidecarContainer</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1PolyaxonSidecarContainer} obj Optional instance to populate.
     * @return {module:model/V1PolyaxonSidecarContainer} The populated <code>V1PolyaxonSidecarContainer</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1PolyaxonSidecarContainer();

            if (data.hasOwnProperty('image')) {
                obj['image'] = ApiClient.convertToType(data['image'], 'String');
            }
            if (data.hasOwnProperty('imageTag')) {
                obj['imageTag'] = ApiClient.convertToType(data['imageTag'], 'String');
            }
            if (data.hasOwnProperty('imagePullPolicy')) {
                obj['imagePullPolicy'] = ApiClient.convertToType(data['imagePullPolicy'], 'String');
            }
            if (data.hasOwnProperty('sleepInterval')) {
                obj['sleepInterval'] = ApiClient.convertToType(data['sleepInterval'], 'Number');
            }
            if (data.hasOwnProperty('syncInterval')) {
                obj['syncInterval'] = ApiClient.convertToType(data['syncInterval'], 'Number');
            }
            if (data.hasOwnProperty('monitorLogs')) {
                obj['monitorLogs'] = ApiClient.convertToType(data['monitorLogs'], 'Boolean');
            }
            if (data.hasOwnProperty('resources')) {
                obj['resources'] = ApiClient.convertToType(data['resources'], Object);
            }
        }
        return obj;
    }


}

/**
 * @member {String} image
 */
V1PolyaxonSidecarContainer.prototype['image'] = undefined;

/**
 * @member {String} imageTag
 */
V1PolyaxonSidecarContainer.prototype['imageTag'] = undefined;

/**
 * @member {String} imagePullPolicy
 */
V1PolyaxonSidecarContainer.prototype['imagePullPolicy'] = undefined;

/**
 * @member {Number} sleepInterval
 */
V1PolyaxonSidecarContainer.prototype['sleepInterval'] = undefined;

/**
 * @member {Number} syncInterval
 */
V1PolyaxonSidecarContainer.prototype['syncInterval'] = undefined;

/**
 * @member {Boolean} monitorLogs
 */
V1PolyaxonSidecarContainer.prototype['monitorLogs'] = undefined;

/**
 * @member {Object} resources
 */
V1PolyaxonSidecarContainer.prototype['resources'] = undefined;






export default V1PolyaxonSidecarContainer;

