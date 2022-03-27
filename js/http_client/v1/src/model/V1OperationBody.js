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
import V1RunPending from './V1RunPending';

/**
 * The V1OperationBody model module.
 * @module model/V1OperationBody
 * @version 1.17.1
 */
class V1OperationBody {
    /**
     * Constructs a new <code>V1OperationBody</code>.
     * @alias module:model/V1OperationBody
     */
    constructor() { 
        
        V1OperationBody.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1OperationBody</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1OperationBody} obj Optional instance to populate.
     * @return {module:model/V1OperationBody} The populated <code>V1OperationBody</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1OperationBody();

            if (data.hasOwnProperty('content')) {
                obj['content'] = ApiClient.convertToType(data['content'], 'String');
            }
            if (data.hasOwnProperty('is_managed')) {
                obj['is_managed'] = ApiClient.convertToType(data['is_managed'], 'Boolean');
            }
            if (data.hasOwnProperty('pending')) {
                obj['pending'] = V1RunPending.constructFromObject(data['pending']);
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], ['String']);
            }
            if (data.hasOwnProperty('meta_info')) {
                obj['meta_info'] = ApiClient.convertToType(data['meta_info'], Object);
            }
        }
        return obj;
    }


}

/**
 * @member {String} content
 */
V1OperationBody.prototype['content'] = undefined;

/**
 * @member {Boolean} is_managed
 */
V1OperationBody.prototype['is_managed'] = undefined;

/**
 * @member {module:model/V1RunPending} pending
 */
V1OperationBody.prototype['pending'] = undefined;

/**
 * @member {String} name
 */
V1OperationBody.prototype['name'] = undefined;

/**
 * @member {String} description
 */
V1OperationBody.prototype['description'] = undefined;

/**
 * @member {Array.<String>} tags
 */
V1OperationBody.prototype['tags'] = undefined;

/**
 * @member {Object} meta_info
 */
V1OperationBody.prototype['meta_info'] = undefined;






export default V1OperationBody;

