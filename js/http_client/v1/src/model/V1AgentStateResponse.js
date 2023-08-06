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
import AgentStateResponseAgentState from './AgentStateResponseAgentState';
import V1Statuses from './V1Statuses';

/**
 * The V1AgentStateResponse model module.
 * @module model/V1AgentStateResponse
 * @version 2.0.0-rc30
 */
class V1AgentStateResponse {
    /**
     * Constructs a new <code>V1AgentStateResponse</code>.
     * @alias module:model/V1AgentStateResponse
     */
    constructor() {

        V1AgentStateResponse.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {
    }

    /**
     * Constructs a <code>V1AgentStateResponse</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1AgentStateResponse} obj Optional instance to populate.
     * @return {module:model/V1AgentStateResponse} The populated <code>V1AgentStateResponse</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1AgentStateResponse();

            if (data.hasOwnProperty('status')) {
                obj['status'] = V1Statuses.constructFromObject(data['status']);
            }
            if (data.hasOwnProperty('state')) {
                obj['state'] = AgentStateResponseAgentState.constructFromObject(data['state']);
            }
            if (data.hasOwnProperty('live_state')) {
                obj['live_state'] = ApiClient.convertToType(data['live_state'], 'Number');
            }
            if (data.hasOwnProperty('compatible_updates')) {
                obj['compatible_updates'] = ApiClient.convertToType(data['compatible_updates'], Object);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1AgentStateResponse</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1AgentStateResponse</code>.
     */
    static validateJSON(data) {
        // validate the optional field `state`
        if (data['state']) { // data not null
          AgentStateResponseAgentState.validateJSON(data['state']);
        }

        return true;
    }


}



/**
 * @member {module:model/V1Statuses} status
 */
V1AgentStateResponse.prototype['status'] = undefined;

/**
 * @member {module:model/AgentStateResponseAgentState} state
 */
V1AgentStateResponse.prototype['state'] = undefined;

/**
 * @member {Number} live_state
 */
V1AgentStateResponse.prototype['live_state'] = undefined;

/**
 * @member {Object} compatible_updates
 */
V1AgentStateResponse.prototype['compatible_updates'] = undefined;






export default V1AgentStateResponse;

