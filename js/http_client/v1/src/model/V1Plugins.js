/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc28
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import V1Notification from './V1Notification';
import V1PolyaxonSidecarContainer from './V1PolyaxonSidecarContainer';

/**
 * The V1Plugins model module.
 * @module model/V1Plugins
 * @version 2.0.0-rc28
 */
class V1Plugins {
    /**
     * Constructs a new <code>V1Plugins</code>.
     * @alias module:model/V1Plugins
     */
    constructor() {

        V1Plugins.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {
    }

    /**
     * Constructs a <code>V1Plugins</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Plugins} obj Optional instance to populate.
     * @return {module:model/V1Plugins} The populated <code>V1Plugins</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Plugins();

            if (data.hasOwnProperty('auth')) {
                obj['auth'] = ApiClient.convertToType(data['auth'], 'Boolean');
            }
            if (data.hasOwnProperty('docker')) {
                obj['docker'] = ApiClient.convertToType(data['docker'], 'Boolean');
            }
            if (data.hasOwnProperty('shm')) {
                obj['shm'] = ApiClient.convertToType(data['shm'], 'Boolean');
            }
            if (data.hasOwnProperty('mountArtifactsStore')) {
                obj['mountArtifactsStore'] = ApiClient.convertToType(data['mountArtifactsStore'], 'Boolean');
            }
            if (data.hasOwnProperty('collectArtifacts')) {
                obj['collectArtifacts'] = ApiClient.convertToType(data['collectArtifacts'], 'Boolean');
            }
            if (data.hasOwnProperty('collectLogs')) {
                obj['collectLogs'] = ApiClient.convertToType(data['collectLogs'], 'Boolean');
            }
            if (data.hasOwnProperty('collectResources')) {
                obj['collectResources'] = ApiClient.convertToType(data['collectResources'], 'Boolean');
            }
            if (data.hasOwnProperty('syncStatuses')) {
                obj['syncStatuses'] = ApiClient.convertToType(data['syncStatuses'], 'Boolean');
            }
            if (data.hasOwnProperty('autoResume')) {
                obj['autoResume'] = ApiClient.convertToType(data['autoResume'], 'Boolean');
            }
            if (data.hasOwnProperty('logLevel')) {
                obj['logLevel'] = ApiClient.convertToType(data['logLevel'], 'String');
            }
            if (data.hasOwnProperty('externalHost')) {
                obj['externalHost'] = ApiClient.convertToType(data['externalHost'], 'Boolean');
            }
            if (data.hasOwnProperty('sidecar')) {
                obj['sidecar'] = V1PolyaxonSidecarContainer.constructFromObject(data['sidecar']);
            }
            if (data.hasOwnProperty('notifications')) {
                obj['notifications'] = ApiClient.convertToType(data['notifications'], [V1Notification]);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1Plugins</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1Plugins</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['logLevel'] && !(typeof data['logLevel'] === 'string' || data['logLevel'] instanceof String)) {
            throw new Error("Expected the field `logLevel` to be a primitive type in the JSON string but got " + data['logLevel']);
        }
        // validate the optional field `sidecar`
        if (data['sidecar']) { // data not null
          V1PolyaxonSidecarContainer.validateJSON(data['sidecar']);
        }
        if (data['notifications']) { // data not null
            // ensure the json data is an array
            if (!Array.isArray(data['notifications'])) {
                throw new Error("Expected the field `notifications` to be an array in the JSON data but got " + data['notifications']);
            }
            // validate the optional field `notifications` (array)
            for (const item of data['notifications']) {
                V1Notification.validateJSON(item);
            };
        }

        return true;
    }


}



/**
 * @member {Boolean} auth
 */
V1Plugins.prototype['auth'] = undefined;

/**
 * @member {Boolean} docker
 */
V1Plugins.prototype['docker'] = undefined;

/**
 * @member {Boolean} shm
 */
V1Plugins.prototype['shm'] = undefined;

/**
 * @member {Boolean} mountArtifactsStore
 */
V1Plugins.prototype['mountArtifactsStore'] = undefined;

/**
 * @member {Boolean} collectArtifacts
 */
V1Plugins.prototype['collectArtifacts'] = undefined;

/**
 * @member {Boolean} collectLogs
 */
V1Plugins.prototype['collectLogs'] = undefined;

/**
 * @member {Boolean} collectResources
 */
V1Plugins.prototype['collectResources'] = undefined;

/**
 * @member {Boolean} syncStatuses
 */
V1Plugins.prototype['syncStatuses'] = undefined;

/**
 * @member {Boolean} autoResume
 */
V1Plugins.prototype['autoResume'] = undefined;

/**
 * @member {String} logLevel
 */
V1Plugins.prototype['logLevel'] = undefined;

/**
 * @member {Boolean} externalHost
 */
V1Plugins.prototype['externalHost'] = undefined;

/**
 * @member {module:model/V1PolyaxonSidecarContainer} sidecar
 */
V1Plugins.prototype['sidecar'] = undefined;

/**
 * @member {Array.<module:model/V1Notification>} notifications
 */
V1Plugins.prototype['notifications'] = undefined;






export default V1Plugins;

