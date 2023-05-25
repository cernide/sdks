/**
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.0.0-rc16
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The V1Environment model module.
 * @module model/V1Environment
 * @version 2.0.0-rc16
 */
class V1Environment {
    /**
     * Constructs a new <code>V1Environment</code>.
     * @alias module:model/V1Environment
     */
    constructor() { 
        
        V1Environment.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1Environment</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Environment} obj Optional instance to populate.
     * @return {module:model/V1Environment} The populated <code>V1Environment</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Environment();

            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], {'String': 'String'});
            }
            if (data.hasOwnProperty('annotations')) {
                obj['annotations'] = ApiClient.convertToType(data['annotations'], {'String': 'String'});
            }
            if (data.hasOwnProperty('nodeSelector')) {
                obj['nodeSelector'] = ApiClient.convertToType(data['nodeSelector'], {'String': 'String'});
            }
            if (data.hasOwnProperty('affinity')) {
                obj['affinity'] = ApiClient.convertToType(data['affinity'], Object);
            }
            if (data.hasOwnProperty('tolerations')) {
                obj['tolerations'] = ApiClient.convertToType(data['tolerations'], [Object]);
            }
            if (data.hasOwnProperty('nodeName')) {
                obj['nodeName'] = ApiClient.convertToType(data['nodeName'], 'String');
            }
            if (data.hasOwnProperty('serviceAccountName')) {
                obj['serviceAccountName'] = ApiClient.convertToType(data['serviceAccountName'], 'String');
            }
            if (data.hasOwnProperty('hostAliases')) {
                obj['hostAliases'] = ApiClient.convertToType(data['hostAliases'], [Object]);
            }
            if (data.hasOwnProperty('securityContext')) {
                obj['securityContext'] = ApiClient.convertToType(data['securityContext'], Object);
            }
            if (data.hasOwnProperty('imagePullSecrets')) {
                obj['imagePullSecrets'] = ApiClient.convertToType(data['imagePullSecrets'], ['String']);
            }
            if (data.hasOwnProperty('hostNetwork')) {
                obj['hostNetwork'] = ApiClient.convertToType(data['hostNetwork'], 'Boolean');
            }
            if (data.hasOwnProperty('hostPID')) {
                obj['hostPID'] = ApiClient.convertToType(data['hostPID'], 'String');
            }
            if (data.hasOwnProperty('dnsPolicy')) {
                obj['dnsPolicy'] = ApiClient.convertToType(data['dnsPolicy'], 'String');
            }
            if (data.hasOwnProperty('dnsConfig')) {
                obj['dnsConfig'] = ApiClient.convertToType(data['dnsConfig'], Object);
            }
            if (data.hasOwnProperty('schedulerName')) {
                obj['schedulerName'] = ApiClient.convertToType(data['schedulerName'], 'String');
            }
            if (data.hasOwnProperty('priorityClassName')) {
                obj['priorityClassName'] = ApiClient.convertToType(data['priorityClassName'], 'String');
            }
            if (data.hasOwnProperty('priority')) {
                obj['priority'] = ApiClient.convertToType(data['priority'], 'Number');
            }
            if (data.hasOwnProperty('restartPolicy')) {
                obj['restartPolicy'] = ApiClient.convertToType(data['restartPolicy'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1Environment</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1Environment</code>.
     */
    static validateJSON(data) {
        // ensure the json data is an array
        if (!Array.isArray(data['tolerations'])) {
            throw new Error("Expected the field `tolerations` to be an array in the JSON data but got " + data['tolerations']);
        }
        // ensure the json data is a string
        if (data['nodeName'] && !(typeof data['nodeName'] === 'string' || data['nodeName'] instanceof String)) {
            throw new Error("Expected the field `nodeName` to be a primitive type in the JSON string but got " + data['nodeName']);
        }
        // ensure the json data is a string
        if (data['serviceAccountName'] && !(typeof data['serviceAccountName'] === 'string' || data['serviceAccountName'] instanceof String)) {
            throw new Error("Expected the field `serviceAccountName` to be a primitive type in the JSON string but got " + data['serviceAccountName']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['hostAliases'])) {
            throw new Error("Expected the field `hostAliases` to be an array in the JSON data but got " + data['hostAliases']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['imagePullSecrets'])) {
            throw new Error("Expected the field `imagePullSecrets` to be an array in the JSON data but got " + data['imagePullSecrets']);
        }
        // ensure the json data is a string
        if (data['hostPID'] && !(typeof data['hostPID'] === 'string' || data['hostPID'] instanceof String)) {
            throw new Error("Expected the field `hostPID` to be a primitive type in the JSON string but got " + data['hostPID']);
        }
        // ensure the json data is a string
        if (data['dnsPolicy'] && !(typeof data['dnsPolicy'] === 'string' || data['dnsPolicy'] instanceof String)) {
            throw new Error("Expected the field `dnsPolicy` to be a primitive type in the JSON string but got " + data['dnsPolicy']);
        }
        // ensure the json data is a string
        if (data['schedulerName'] && !(typeof data['schedulerName'] === 'string' || data['schedulerName'] instanceof String)) {
            throw new Error("Expected the field `schedulerName` to be a primitive type in the JSON string but got " + data['schedulerName']);
        }
        // ensure the json data is a string
        if (data['priorityClassName'] && !(typeof data['priorityClassName'] === 'string' || data['priorityClassName'] instanceof String)) {
            throw new Error("Expected the field `priorityClassName` to be a primitive type in the JSON string but got " + data['priorityClassName']);
        }
        // ensure the json data is a string
        if (data['restartPolicy'] && !(typeof data['restartPolicy'] === 'string' || data['restartPolicy'] instanceof String)) {
            throw new Error("Expected the field `restartPolicy` to be a primitive type in the JSON string but got " + data['restartPolicy']);
        }

        return true;
    }


}



/**
 * @member {Object.<String, String>} labels
 */
V1Environment.prototype['labels'] = undefined;

/**
 * @member {Object.<String, String>} annotations
 */
V1Environment.prototype['annotations'] = undefined;

/**
 * @member {Object.<String, String>} nodeSelector
 */
V1Environment.prototype['nodeSelector'] = undefined;

/**
 * @member {Object} affinity
 */
V1Environment.prototype['affinity'] = undefined;

/**
 * Optional Tolerations to apply.
 * @member {Array.<Object>} tolerations
 */
V1Environment.prototype['tolerations'] = undefined;

/**
 * Optional NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.
 * @member {String} nodeName
 */
V1Environment.prototype['nodeName'] = undefined;

/**
 * @member {String} serviceAccountName
 */
V1Environment.prototype['serviceAccountName'] = undefined;

/**
 * Optional HostAliases is an optional list of hosts and IPs that will be injected into the pod spec.
 * @member {Array.<Object>} hostAliases
 */
V1Environment.prototype['hostAliases'] = undefined;

/**
 * @member {Object} securityContext
 */
V1Environment.prototype['securityContext'] = undefined;

/**
 * @member {Array.<String>} imagePullSecrets
 */
V1Environment.prototype['imagePullSecrets'] = undefined;

/**
 * Host networking requested for this workflow pod. Default to false.
 * @member {Boolean} hostNetwork
 */
V1Environment.prototype['hostNetwork'] = undefined;

/**
 * Use the host's pid namespace. Default to false.
 * @member {String} hostPID
 */
V1Environment.prototype['hostPID'] = undefined;

/**
 * Set DNS policy for the pod. Defaults to \"ClusterFirst\". Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'.
 * @member {String} dnsPolicy
 */
V1Environment.prototype['dnsPolicy'] = undefined;

/**
 * @member {Object} dnsConfig
 */
V1Environment.prototype['dnsConfig'] = undefined;

/**
 * @member {String} schedulerName
 */
V1Environment.prototype['schedulerName'] = undefined;

/**
 * If specified, indicates the pod's priority. \"system-node-critical\" and \"system-cluster-critical\" are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.
 * @member {String} priorityClassName
 */
V1Environment.prototype['priorityClassName'] = undefined;

/**
 * The priority value. Various system components use this field to find the priority of the pod. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority.
 * @member {Number} priority
 */
V1Environment.prototype['priority'] = undefined;

/**
 * @member {String} restartPolicy
 */
V1Environment.prototype['restartPolicy'] = undefined;






export default V1Environment;

