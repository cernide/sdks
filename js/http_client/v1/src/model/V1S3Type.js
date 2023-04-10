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
 * The version of the OpenAPI document: 2.0.0-rc2
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The V1S3Type model module.
 * @module model/V1S3Type
 * @version 2.0.0-rc2
 */
class V1S3Type {
    /**
     * Constructs a new <code>V1S3Type</code>.
     * @alias module:model/V1S3Type
     */
    constructor() {

        V1S3Type.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {
    }

    /**
     * Constructs a <code>V1S3Type</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1S3Type} obj Optional instance to populate.
     * @return {module:model/V1S3Type} The populated <code>V1S3Type</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1S3Type();

            if (data.hasOwnProperty('bucket')) {
                obj['bucket'] = ApiClient.convertToType(data['bucket'], 'String');
            }
            if (data.hasOwnProperty('key')) {
                obj['key'] = ApiClient.convertToType(data['key'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1S3Type</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1S3Type</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['bucket'] && !(typeof data['bucket'] === 'string' || data['bucket'] instanceof String)) {
            throw new Error("Expected the field `bucket` to be a primitive type in the JSON string but got " + data['bucket']);
        }
        // ensure the json data is a string
        if (data['key'] && !(typeof data['key'] === 'string' || data['key'] instanceof String)) {
            throw new Error("Expected the field `key` to be a primitive type in the JSON string but got " + data['key']);
        }

        return true;
    }


}



/**
 * @member {String} bucket
 */
V1S3Type.prototype['bucket'] = undefined;

/**
 * @member {String} key
 */
V1S3Type.prototype['key'] = undefined;






export default V1S3Type;

