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

/**
 * The V1Validation model module.
 * @module model/V1Validation
 * @version 2.4.2
 */
class V1Validation {
    /**
     * Constructs a new <code>V1Validation</code>.
     * @alias module:model/V1Validation
     */
    constructor() {
        V1Validation.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {}

    /**
     * Constructs a <code>V1Validation</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Validation} obj Optional instance to populate.
     * @return {module:model/V1Validation} The populated <code>V1Validation</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Validation();

            if (data.hasOwnProperty("delay")) {
                obj["delay"] = ApiClient.convertToType(
                    data["delay"],
                    "Boolean"
                );
            }
            if (data.hasOwnProperty("gt")) {
                obj["gt"] = ApiClient.convertToType(data["gt"], "Number");
            }
            if (data.hasOwnProperty("ge")) {
                obj["ge"] = ApiClient.convertToType(data["ge"], "Number");
            }
            if (data.hasOwnProperty("lt")) {
                obj["lt"] = ApiClient.convertToType(data["lt"], "Number");
            }
            if (data.hasOwnProperty("le")) {
                obj["le"] = ApiClient.convertToType(data["le"], "Number");
            }
            if (data.hasOwnProperty("multipleOf")) {
                obj["multipleOf"] = ApiClient.convertToType(
                    data["multipleOf"],
                    "Number"
                );
            }
            if (data.hasOwnProperty("minDigits")) {
                obj["minDigits"] = ApiClient.convertToType(
                    data["minDigits"],
                    "Number"
                );
            }
            if (data.hasOwnProperty("maxDigits")) {
                obj["maxDigits"] = ApiClient.convertToType(
                    data["maxDigits"],
                    "Number"
                );
            }
            if (data.hasOwnProperty("decimalPlaces")) {
                obj["decimalPlaces"] = ApiClient.convertToType(
                    data["decimalPlaces"],
                    "Number"
                );
            }
            if (data.hasOwnProperty("regex")) {
                obj["regex"] = ApiClient.convertToType(data["regex"], "String");
            }
            if (data.hasOwnProperty("minLength")) {
                obj["minLength"] = ApiClient.convertToType(
                    data["minLength"],
                    "Number"
                );
            }
            if (data.hasOwnProperty("maxLength")) {
                obj["maxLength"] = ApiClient.convertToType(
                    data["maxLength"],
                    "Number"
                );
            }
            if (data.hasOwnProperty("contains")) {
                obj["contains"] = ApiClient.convertToType(data["contains"], [
                    Object,
                ]);
            }
            if (data.hasOwnProperty("excludes")) {
                obj["excludes"] = ApiClient.convertToType(data["excludes"], [
                    Object,
                ]);
            }
            if (data.hasOwnProperty("options")) {
                obj["options"] = ApiClient.convertToType(data["options"], [
                    Object,
                ]);
            }
            if (data.hasOwnProperty("minItems")) {
                obj["minItems"] = ApiClient.convertToType(
                    data["minItems"],
                    "Number"
                );
            }
            if (data.hasOwnProperty("maxItems")) {
                obj["maxItems"] = ApiClient.convertToType(
                    data["maxItems"],
                    "Number"
                );
            }
            if (data.hasOwnProperty("keys")) {
                obj["keys"] = ApiClient.convertToType(data["keys"], ["String"]);
            }
            if (data.hasOwnProperty("containsKeys")) {
                obj["containsKeys"] = ApiClient.convertToType(
                    data["containsKeys"],
                    ["String"]
                );
            }
            if (data.hasOwnProperty("excludesKeys")) {
                obj["excludesKeys"] = ApiClient.convertToType(
                    data["excludesKeys"],
                    ["String"]
                );
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1Validation</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1Validation</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (
            data["regex"] &&
            !(
                typeof data["regex"] === "string" ||
                data["regex"] instanceof String
            )
        ) {
            throw new Error(
                "Expected the field `regex` to be a primitive type in the JSON string but got " +
                    data["regex"]
            );
        }
        // ensure the json data is an array
        if (!Array.isArray(data["contains"])) {
            throw new Error(
                "Expected the field `contains` to be an array in the JSON data but got " +
                    data["contains"]
            );
        }
        // ensure the json data is an array
        if (!Array.isArray(data["excludes"])) {
            throw new Error(
                "Expected the field `excludes` to be an array in the JSON data but got " +
                    data["excludes"]
            );
        }
        // ensure the json data is an array
        if (!Array.isArray(data["options"])) {
            throw new Error(
                "Expected the field `options` to be an array in the JSON data but got " +
                    data["options"]
            );
        }
        // ensure the json data is an array
        if (!Array.isArray(data["keys"])) {
            throw new Error(
                "Expected the field `keys` to be an array in the JSON data but got " +
                    data["keys"]
            );
        }
        // ensure the json data is an array
        if (!Array.isArray(data["containsKeys"])) {
            throw new Error(
                "Expected the field `containsKeys` to be an array in the JSON data but got " +
                    data["containsKeys"]
            );
        }
        // ensure the json data is an array
        if (!Array.isArray(data["excludesKeys"])) {
            throw new Error(
                "Expected the field `excludesKeys` to be an array in the JSON data but got " +
                    data["excludesKeys"]
            );
        }

        return true;
    }
}

/**
 * @member {Boolean} delay
 */
V1Validation.prototype["delay"] = undefined;

/**
 * @member {Number} gt
 */
V1Validation.prototype["gt"] = undefined;

/**
 * @member {Number} ge
 */
V1Validation.prototype["ge"] = undefined;

/**
 * @member {Number} lt
 */
V1Validation.prototype["lt"] = undefined;

/**
 * @member {Number} le
 */
V1Validation.prototype["le"] = undefined;

/**
 * @member {Number} multipleOf
 */
V1Validation.prototype["multipleOf"] = undefined;

/**
 * @member {Number} minDigits
 */
V1Validation.prototype["minDigits"] = undefined;

/**
 * @member {Number} maxDigits
 */
V1Validation.prototype["maxDigits"] = undefined;

/**
 * @member {Number} decimalPlaces
 */
V1Validation.prototype["decimalPlaces"] = undefined;

/**
 * @member {String} regex
 */
V1Validation.prototype["regex"] = undefined;

/**
 * @member {Number} minLength
 */
V1Validation.prototype["minLength"] = undefined;

/**
 * @member {Number} maxLength
 */
V1Validation.prototype["maxLength"] = undefined;

/**
 * @member {Array.<Object>} contains
 */
V1Validation.prototype["contains"] = undefined;

/**
 * @member {Array.<Object>} excludes
 */
V1Validation.prototype["excludes"] = undefined;

/**
 * @member {Array.<Object>} options
 */
V1Validation.prototype["options"] = undefined;

/**
 * @member {Number} minItems
 */
V1Validation.prototype["minItems"] = undefined;

/**
 * @member {Number} maxItems
 */
V1Validation.prototype["maxItems"] = undefined;

/**
 * @member {Array.<String>} keys
 */
V1Validation.prototype["keys"] = undefined;

/**
 * @member {Array.<String>} containsKeys
 */
V1Validation.prototype["containsKeys"] = undefined;

/**
 * @member {Array.<String>} excludesKeys
 */
V1Validation.prototype["excludesKeys"] = undefined;

export default V1Validation;
