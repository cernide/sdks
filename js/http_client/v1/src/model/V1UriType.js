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
 * The version of the OpenAPI document: 2.0.0-rc10
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The V1UriType model module.
 * @module model/V1UriType
 * @version 2.0.0-rc10
 */
class V1UriType {
    /**
     * Constructs a new <code>V1UriType</code>.
     * @alias module:model/V1UriType
     */
    constructor() {

        V1UriType.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {
    }

    /**
     * Constructs a <code>V1UriType</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1UriType} obj Optional instance to populate.
     * @return {module:model/V1UriType} The populated <code>V1UriType</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1UriType();

            if (data.hasOwnProperty('user')) {
                obj['user'] = ApiClient.convertToType(data['user'], 'String');
            }
            if (data.hasOwnProperty('password')) {
                obj['password'] = ApiClient.convertToType(data['password'], 'String');
            }
            if (data.hasOwnProperty('host')) {
                obj['host'] = ApiClient.convertToType(data['host'], 'Boolean');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1UriType</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1UriType</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['user'] && !(typeof data['user'] === 'string' || data['user'] instanceof String)) {
            throw new Error("Expected the field `user` to be a primitive type in the JSON string but got " + data['user']);
        }
        // ensure the json data is a string
        if (data['password'] && !(typeof data['password'] === 'string' || data['password'] instanceof String)) {
            throw new Error("Expected the field `password` to be a primitive type in the JSON string but got " + data['password']);
        }

        return true;
    }


}



/**
 * @member {String} user
 */
V1UriType.prototype['user'] = undefined;

/**
 * @member {String} password
 */
V1UriType.prototype['password'] = undefined;

/**
 * @member {Boolean} host
 */
V1UriType.prototype['host'] = undefined;






export default V1UriType;

