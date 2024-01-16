/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.1.0-rc5
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The V1GitConnection model module.
 * @module model/V1GitConnection
 * @version 2.1.0-rc5
 */
class V1GitConnection {
    /**
     * Constructs a new <code>V1GitConnection</code>.
     * @alias module:model/V1GitConnection
     */
    constructor() {

        V1GitConnection.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {
    }

    /**
     * Constructs a <code>V1GitConnection</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1GitConnection} obj Optional instance to populate.
     * @return {module:model/V1GitConnection} The populated <code>V1GitConnection</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1GitConnection();

            if (data.hasOwnProperty('url')) {
                obj['url'] = ApiClient.convertToType(data['url'], 'String');
            }
            if (data.hasOwnProperty('revision')) {
                obj['revision'] = ApiClient.convertToType(data['revision'], 'Boolean');
            }
            if (data.hasOwnProperty('flags')) {
                obj['flags'] = ApiClient.convertToType(data['flags'], ['String']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1GitConnection</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1GitConnection</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['url'] && !(typeof data['url'] === 'string' || data['url'] instanceof String)) {
            throw new Error("Expected the field `url` to be a primitive type in the JSON string but got " + data['url']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['flags'])) {
            throw new Error("Expected the field `flags` to be an array in the JSON data but got " + data['flags']);
        }

        return true;
    }


}



/**
 * @member {String} url
 */
V1GitConnection.prototype['url'] = undefined;

/**
 * @member {Boolean} revision
 */
V1GitConnection.prototype['revision'] = undefined;

/**
 * @member {Array.<String>} flags
 */
V1GitConnection.prototype['flags'] = undefined;






export default V1GitConnection;

