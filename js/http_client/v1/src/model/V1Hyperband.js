/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.4.2
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from "../ApiClient";
import V1OptimizationMetric from "./V1OptimizationMetric";
import V1OptimizationResource from "./V1OptimizationResource";
import V1Tuner from "./V1Tuner";

/**
 * The V1Hyperband model module.
 * @module model/V1Hyperband
 * @version 2.4.2
 */
class V1Hyperband {
    /**
     * Constructs a new <code>V1Hyperband</code>.
     * @alias module:model/V1Hyperband
     */
    constructor() {
        V1Hyperband.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {}

    /**
     * Constructs a <code>V1Hyperband</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Hyperband} obj Optional instance to populate.
     * @return {module:model/V1Hyperband} The populated <code>V1Hyperband</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Hyperband();

            if (data.hasOwnProperty("kind")) {
                obj["kind"] = ApiClient.convertToType(data["kind"], "String");
            }
            if (data.hasOwnProperty("params")) {
                obj["params"] = ApiClient.convertToType(data["params"], {
                    String: Object,
                });
            }
            if (data.hasOwnProperty("maxIterations")) {
                obj["maxIterations"] = ApiClient.convertToType(
                    data["maxIterations"],
                    "Number"
                );
            }
            if (data.hasOwnProperty("eta")) {
                obj["eta"] = ApiClient.convertToType(data["eta"], "Number");
            }
            if (data.hasOwnProperty("resource")) {
                obj["resource"] = V1OptimizationResource.constructFromObject(
                    data["resource"]
                );
            }
            if (data.hasOwnProperty("metric")) {
                obj["metric"] = V1OptimizationMetric.constructFromObject(
                    data["metric"]
                );
            }
            if (data.hasOwnProperty("resume")) {
                obj["resume"] = ApiClient.convertToType(
                    data["resume"],
                    "Boolean"
                );
            }
            if (data.hasOwnProperty("seed")) {
                obj["seed"] = ApiClient.convertToType(data["seed"], "Number");
            }
            if (data.hasOwnProperty("concurrency")) {
                obj["concurrency"] = ApiClient.convertToType(
                    data["concurrency"],
                    "Number"
                );
            }
            if (data.hasOwnProperty("tuner")) {
                obj["tuner"] = V1Tuner.constructFromObject(data["tuner"]);
            }
            if (data.hasOwnProperty("earlyStopping")) {
                obj["earlyStopping"] = ApiClient.convertToType(
                    data["earlyStopping"],
                    [Object]
                );
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1Hyperband</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1Hyperband</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (
            data["kind"] &&
            !(
                typeof data["kind"] === "string" ||
                data["kind"] instanceof String
            )
        ) {
            throw new Error(
                "Expected the field `kind` to be a primitive type in the JSON string but got " +
                    data["kind"]
            );
        }
        // validate the optional field `resource`
        if (data["resource"]) {
            // data not null
            V1OptimizationResource.validateJSON(data["resource"]);
        }
        // validate the optional field `metric`
        if (data["metric"]) {
            // data not null
            V1OptimizationMetric.validateJSON(data["metric"]);
        }
        // validate the optional field `tuner`
        if (data["tuner"]) {
            // data not null
            V1Tuner.validateJSON(data["tuner"]);
        }
        // ensure the json data is an array
        if (!Array.isArray(data["earlyStopping"])) {
            throw new Error(
                "Expected the field `earlyStopping` to be an array in the JSON data but got " +
                    data["earlyStopping"]
            );
        }

        return true;
    }
}

/**
 * @member {String} kind
 * @default 'hyperband'
 */
V1Hyperband.prototype["kind"] = "hyperband";

/**
 * @member {Object.<String, Object>} params
 */
V1Hyperband.prototype["params"] = undefined;

/**
 * @member {Number} maxIterations
 */
V1Hyperband.prototype["maxIterations"] = undefined;

/**
 * @member {Number} eta
 */
V1Hyperband.prototype["eta"] = undefined;

/**
 * @member {module:model/V1OptimizationResource} resource
 */
V1Hyperband.prototype["resource"] = undefined;

/**
 * @member {module:model/V1OptimizationMetric} metric
 */
V1Hyperband.prototype["metric"] = undefined;

/**
 * @member {Boolean} resume
 */
V1Hyperband.prototype["resume"] = undefined;

/**
 * @member {Number} seed
 */
V1Hyperband.prototype["seed"] = undefined;

/**
 * @member {Number} concurrency
 */
V1Hyperband.prototype["concurrency"] = undefined;

/**
 * @member {module:model/V1Tuner} tuner
 */
V1Hyperband.prototype["tuner"] = undefined;

/**
 * @member {Array.<Object>} earlyStopping
 */
V1Hyperband.prototype["earlyStopping"] = undefined;

export default V1Hyperband;
