/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The V1GcsType model module.
 * @module model/V1GcsType
 * @version 2.0.0
 */
class V1GcsType {
    /**
     * Constructs a new <code>V1GcsType</code>.
     * @alias module:model/V1GcsType
     */
    constructor() {

        V1GcsType.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {
    }

    /**
     * Constructs a <code>V1GcsType</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1GcsType} obj Optional instance to populate.
     * @return {module:model/V1GcsType} The populated <code>V1GcsType</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1GcsType();

            if (data.hasOwnProperty('bucket')) {
                obj['bucket'] = ApiClient.convertToType(data['bucket'], 'String');
            }
            if (data.hasOwnProperty('blob')) {
                obj['blob'] = ApiClient.convertToType(data['blob'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1GcsType</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1GcsType</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['bucket'] && !(typeof data['bucket'] === 'string' || data['bucket'] instanceof String)) {
            throw new Error("Expected the field `bucket` to be a primitive type in the JSON string but got " + data['bucket']);
        }
        // ensure the json data is a string
        if (data['blob'] && !(typeof data['blob'] === 'string' || data['blob'] instanceof String)) {
            throw new Error("Expected the field `blob` to be a primitive type in the JSON string but got " + data['blob']);
        }

        return true;
    }


}



/**
 * @member {String} bucket
 */
V1GcsType.prototype['bucket'] = undefined;

/**
 * @member {String} blob
 */
V1GcsType.prototype['blob'] = undefined;






export default V1GcsType;

