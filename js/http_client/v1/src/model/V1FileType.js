/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc19
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The V1FileType model module.
 * @module model/V1FileType
 * @version 2.0.0-rc19
 */
class V1FileType {
    /**
     * Constructs a new <code>V1FileType</code>.
     * @alias module:model/V1FileType
     */
    constructor() {

        V1FileType.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {
    }

    /**
     * Constructs a <code>V1FileType</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1FileType} obj Optional instance to populate.
     * @return {module:model/V1FileType} The populated <code>V1FileType</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1FileType();

            if (data.hasOwnProperty('content')) {
                obj['content'] = ApiClient.convertToType(data['content'], 'String');
            }
            if (data.hasOwnProperty('filename')) {
                obj['filename'] = ApiClient.convertToType(data['filename'], 'String');
            }
            if (data.hasOwnProperty('chmod')) {
                obj['chmod'] = ApiClient.convertToType(data['chmod'], 'String');
            }
            if (data.hasOwnProperty('kind')) {
                obj['kind'] = ApiClient.convertToType(data['kind'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1FileType</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1FileType</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['content'] && !(typeof data['content'] === 'string' || data['content'] instanceof String)) {
            throw new Error("Expected the field `content` to be a primitive type in the JSON string but got " + data['content']);
        }
        // ensure the json data is a string
        if (data['filename'] && !(typeof data['filename'] === 'string' || data['filename'] instanceof String)) {
            throw new Error("Expected the field `filename` to be a primitive type in the JSON string but got " + data['filename']);
        }
        // ensure the json data is a string
        if (data['chmod'] && !(typeof data['chmod'] === 'string' || data['chmod'] instanceof String)) {
            throw new Error("Expected the field `chmod` to be a primitive type in the JSON string but got " + data['chmod']);
        }
        // ensure the json data is a string
        if (data['kind'] && !(typeof data['kind'] === 'string' || data['kind'] instanceof String)) {
            throw new Error("Expected the field `kind` to be a primitive type in the JSON string but got " + data['kind']);
        }

        return true;
    }


}



/**
 * @member {String} content
 */
V1FileType.prototype['content'] = undefined;

/**
 * @member {String} filename
 */
V1FileType.prototype['filename'] = undefined;

/**
 * @member {String} chmod
 */
V1FileType.prototype['chmod'] = undefined;

/**
 * @member {String} kind
 */
V1FileType.prototype['kind'] = undefined;






export default V1FileType;

