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
 * The version of the OpenAPI document: 1.13.0-rc2
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import ProtobufAny from './ProtobufAny';

/**
 * The RuntimeError model module.
 * @module model/RuntimeError
 * @version 1.13.0-rc2
 */
class RuntimeError {
    /**
     * Constructs a new <code>RuntimeError</code>.
     * @alias module:model/RuntimeError
     */
    constructor() { 
        
        RuntimeError.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>RuntimeError</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/RuntimeError} obj Optional instance to populate.
     * @return {module:model/RuntimeError} The populated <code>RuntimeError</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RuntimeError();

            if (data.hasOwnProperty('error')) {
                obj['error'] = ApiClient.convertToType(data['error'], 'String');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'Number');
            }
            if (data.hasOwnProperty('message')) {
                obj['message'] = ApiClient.convertToType(data['message'], 'String');
            }
            if (data.hasOwnProperty('details')) {
                obj['details'] = ApiClient.convertToType(data['details'], [ProtobufAny]);
            }
        }
        return obj;
    }


}

/**
 * @member {String} error
 */
RuntimeError.prototype['error'] = undefined;

/**
 * @member {Number} code
 */
RuntimeError.prototype['code'] = undefined;

/**
 * @member {String} message
 */
RuntimeError.prototype['message'] = undefined;

/**
 * @member {Array.<module:model/ProtobufAny>} details
 */
RuntimeError.prototype['details'] = undefined;






export default RuntimeError;

