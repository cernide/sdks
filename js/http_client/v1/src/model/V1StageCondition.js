/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc30
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import V1Stages from './V1Stages';

/**
 * The V1StageCondition model module.
 * @module model/V1StageCondition
 * @version 2.0.0-rc30
 */
class V1StageCondition {
    /**
     * Constructs a new <code>V1StageCondition</code>.
     * @alias module:model/V1StageCondition
     */
    constructor() {

        V1StageCondition.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {
    }

    /**
     * Constructs a <code>V1StageCondition</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1StageCondition} obj Optional instance to populate.
     * @return {module:model/V1StageCondition} The populated <code>V1StageCondition</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1StageCondition();

            if (data.hasOwnProperty('type')) {
                obj['type'] = V1Stages.constructFromObject(data['type']);
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
            if (data.hasOwnProperty('reason')) {
                obj['reason'] = ApiClient.convertToType(data['reason'], 'String');
            }
            if (data.hasOwnProperty('message')) {
                obj['message'] = ApiClient.convertToType(data['message'], 'String');
            }
            if (data.hasOwnProperty('last_update_time')) {
                obj['last_update_time'] = ApiClient.convertToType(data['last_update_time'], 'Date');
            }
            if (data.hasOwnProperty('last_transition_time')) {
                obj['last_transition_time'] = ApiClient.convertToType(data['last_transition_time'], 'Date');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1StageCondition</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1StageCondition</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['status'] && !(typeof data['status'] === 'string' || data['status'] instanceof String)) {
            throw new Error("Expected the field `status` to be a primitive type in the JSON string but got " + data['status']);
        }
        // ensure the json data is a string
        if (data['reason'] && !(typeof data['reason'] === 'string' || data['reason'] instanceof String)) {
            throw new Error("Expected the field `reason` to be a primitive type in the JSON string but got " + data['reason']);
        }
        // ensure the json data is a string
        if (data['message'] && !(typeof data['message'] === 'string' || data['message'] instanceof String)) {
            throw new Error("Expected the field `message` to be a primitive type in the JSON string but got " + data['message']);
        }

        return true;
    }


}



/**
 * @member {module:model/V1Stages} type
 */
V1StageCondition.prototype['type'] = undefined;

/**
 * @member {String} status
 */
V1StageCondition.prototype['status'] = undefined;

/**
 * @member {String} reason
 */
V1StageCondition.prototype['reason'] = undefined;

/**
 * @member {String} message
 */
V1StageCondition.prototype['message'] = undefined;

/**
 * @member {Date} last_update_time
 */
V1StageCondition.prototype['last_update_time'] = undefined;

/**
 * @member {Date} last_transition_time
 */
V1StageCondition.prototype['last_transition_time'] = undefined;






export default V1StageCondition;

