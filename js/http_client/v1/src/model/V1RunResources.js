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
 * The V1RunResources model module.
 * @module model/V1RunResources
 * @version 2.0.0
 */
class V1RunResources {
    /**
     * Constructs a new <code>V1RunResources</code>.
     * @alias module:model/V1RunResources
     */
    constructor() {

        V1RunResources.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {
    }

    /**
     * Constructs a <code>V1RunResources</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1RunResources} obj Optional instance to populate.
     * @return {module:model/V1RunResources} The populated <code>V1RunResources</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1RunResources();

            if (data.hasOwnProperty('cpu')) {
                obj['cpu'] = ApiClient.convertToType(data['cpu'], 'Number');
            }
            if (data.hasOwnProperty('memory')) {
                obj['memory'] = ApiClient.convertToType(data['memory'], 'Number');
            }
            if (data.hasOwnProperty('gpu')) {
                obj['gpu'] = ApiClient.convertToType(data['gpu'], 'Number');
            }
            if (data.hasOwnProperty('custom')) {
                obj['custom'] = ApiClient.convertToType(data['custom'], 'Number');
            }
            if (data.hasOwnProperty('cost')) {
                obj['cost'] = ApiClient.convertToType(data['cost'], 'Number');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1RunResources</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1RunResources</code>.
     */
    static validateJSON(data) {

        return true;
    }


}



/**
 * @member {Number} cpu
 */
V1RunResources.prototype['cpu'] = undefined;

/**
 * @member {Number} memory
 */
V1RunResources.prototype['memory'] = undefined;

/**
 * @member {Number} gpu
 */
V1RunResources.prototype['gpu'] = undefined;

/**
 * @member {Number} custom
 */
V1RunResources.prototype['custom'] = undefined;

/**
 * @member {Number} cost
 */
V1RunResources.prototype['cost'] = undefined;






export default V1RunResources;

