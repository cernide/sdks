// Copyright 2018-2023 Polyaxon, Inc.
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
 * The version of the OpenAPI document: 2.0.0-rc5
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import V1ConnectionKind from './V1ConnectionKind';

/**
 * The V1ConnectionResponse model module.
 * @module model/V1ConnectionResponse
 * @version 2.0.0-rc5
 */
class V1ConnectionResponse {
    /**
     * Constructs a new <code>V1ConnectionResponse</code>.
     * @alias module:model/V1ConnectionResponse
     */
    constructor() {

        V1ConnectionResponse.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {
    }

    /**
     * Constructs a <code>V1ConnectionResponse</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1ConnectionResponse} obj Optional instance to populate.
     * @return {module:model/V1ConnectionResponse} The populated <code>V1ConnectionResponse</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1ConnectionResponse();

            if (data.hasOwnProperty('uuid')) {
                obj['uuid'] = ApiClient.convertToType(data['uuid'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('agent')) {
                obj['agent'] = ApiClient.convertToType(data['agent'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], ['String']);
            }
            if (data.hasOwnProperty('created_at')) {
                obj['created_at'] = ApiClient.convertToType(data['created_at'], 'Date');
            }
            if (data.hasOwnProperty('updated_at')) {
                obj['updated_at'] = ApiClient.convertToType(data['updated_at'], 'Date');
            }
            if (data.hasOwnProperty('live_state')) {
                obj['live_state'] = ApiClient.convertToType(data['live_state'], 'Number');
            }
            if (data.hasOwnProperty('kind')) {
                obj['kind'] = V1ConnectionKind.constructFromObject(data['kind']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1ConnectionResponse</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1ConnectionResponse</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['uuid'] && !(typeof data['uuid'] === 'string' || data['uuid'] instanceof String)) {
            throw new Error("Expected the field `uuid` to be a primitive type in the JSON string but got " + data['uuid']);
        }
        // ensure the json data is a string
        if (data['name'] && !(typeof data['name'] === 'string' || data['name'] instanceof String)) {
            throw new Error("Expected the field `name` to be a primitive type in the JSON string but got " + data['name']);
        }
        // ensure the json data is a string
        if (data['agent'] && !(typeof data['agent'] === 'string' || data['agent'] instanceof String)) {
            throw new Error("Expected the field `agent` to be a primitive type in the JSON string but got " + data['agent']);
        }
        // ensure the json data is a string
        if (data['description'] && !(typeof data['description'] === 'string' || data['description'] instanceof String)) {
            throw new Error("Expected the field `description` to be a primitive type in the JSON string but got " + data['description']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['tags'])) {
            throw new Error("Expected the field `tags` to be an array in the JSON data but got " + data['tags']);
        }

        return true;
    }


}



/**
 * @member {String} uuid
 */
V1ConnectionResponse.prototype['uuid'] = undefined;

/**
 * @member {String} name
 */
V1ConnectionResponse.prototype['name'] = undefined;

/**
 * @member {String} agent
 */
V1ConnectionResponse.prototype['agent'] = undefined;

/**
 * @member {String} description
 */
V1ConnectionResponse.prototype['description'] = undefined;

/**
 * @member {Array.<String>} tags
 */
V1ConnectionResponse.prototype['tags'] = undefined;

/**
 * @member {Date} created_at
 */
V1ConnectionResponse.prototype['created_at'] = undefined;

/**
 * @member {Date} updated_at
 */
V1ConnectionResponse.prototype['updated_at'] = undefined;

/**
 * @member {Number} live_state
 */
V1ConnectionResponse.prototype['live_state'] = undefined;

/**
 * @member {module:model/V1ConnectionKind} kind
 */
V1ConnectionResponse.prototype['kind'] = undefined;






export default V1ConnectionResponse;

