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
 * The V1Version model module.
 * @module model/V1Version
 * @version 2.1.0-rc5
 */
class V1Version {
    /**
     * Constructs a new <code>V1Version</code>.
     * @alias module:model/V1Version
     */
    constructor() {

        V1Version.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {
    }

    /**
     * Constructs a <code>V1Version</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Version} obj Optional instance to populate.
     * @return {module:model/V1Version} The populated <code>V1Version</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Version();

            if (data.hasOwnProperty('min')) {
                obj['min'] = ApiClient.convertToType(data['min'], 'String');
            }
            if (data.hasOwnProperty('latest')) {
                obj['latest'] = ApiClient.convertToType(data['latest'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1Version</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1Version</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['min'] && !(typeof data['min'] === 'string' || data['min'] instanceof String)) {
            throw new Error("Expected the field `min` to be a primitive type in the JSON string but got " + data['min']);
        }
        // ensure the json data is a string
        if (data['latest'] && !(typeof data['latest'] === 'string' || data['latest'] instanceof String)) {
            throw new Error("Expected the field `latest` to be a primitive type in the JSON string but got " + data['latest']);
        }

        return true;
    }


}



/**
 * @member {String} min
 */
V1Version.prototype['min'] = undefined;

/**
 * @member {String} latest
 */
V1Version.prototype['latest'] = undefined;






export default V1Version;

