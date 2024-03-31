/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.1.4
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The V1LogHandler model module.
 * @module model/V1LogHandler
 * @version 2.1.4
 */
class V1LogHandler {
    /**
     * Constructs a new <code>V1LogHandler</code>.
     * @alias module:model/V1LogHandler
     */
    constructor() {

        V1LogHandler.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {
    }

    /**
     * Constructs a <code>V1LogHandler</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1LogHandler} obj Optional instance to populate.
     * @return {module:model/V1LogHandler} The populated <code>V1LogHandler</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1LogHandler();

            if (data.hasOwnProperty('dsn')) {
                obj['dsn'] = ApiClient.convertToType(data['dsn'], 'String');
            }
            if (data.hasOwnProperty('environment')) {
                obj['environment'] = ApiClient.convertToType(data['environment'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1LogHandler</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1LogHandler</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['dsn'] && !(typeof data['dsn'] === 'string' || data['dsn'] instanceof String)) {
            throw new Error("Expected the field `dsn` to be a primitive type in the JSON string but got " + data['dsn']);
        }
        // ensure the json data is a string
        if (data['environment'] && !(typeof data['environment'] === 'string' || data['environment'] instanceof String)) {
            throw new Error("Expected the field `environment` to be a primitive type in the JSON string but got " + data['environment']);
        }

        return true;
    }


}



/**
 * @member {String} dsn
 */
V1LogHandler.prototype['dsn'] = undefined;

/**
 * @member {String} environment
 */
V1LogHandler.prototype['environment'] = undefined;






export default V1LogHandler;

