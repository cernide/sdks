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
import V1ArtifactKind from "./V1ArtifactKind";

/**
 * The V1EventType model module.
 * @module model/V1EventType
 * @version 2.4.2
 */
class V1EventType {
    /**
     * Constructs a new <code>V1EventType</code>.
     * @alias module:model/V1EventType
     */
    constructor() {
        V1EventType.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {}

    /**
     * Constructs a <code>V1EventType</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1EventType} obj Optional instance to populate.
     * @return {module:model/V1EventType} The populated <code>V1EventType</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1EventType();

            if (data.hasOwnProperty("name")) {
                obj["name"] = ApiClient.convertToType(data["name"], "String");
            }
            if (data.hasOwnProperty("kind")) {
                obj["kind"] = V1ArtifactKind.constructFromObject(data["kind"]);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1EventType</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1EventType</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (
            data["name"] &&
            !(
                typeof data["name"] === "string" ||
                data["name"] instanceof String
            )
        ) {
            throw new Error(
                "Expected the field `name` to be a primitive type in the JSON string but got " +
                    data["name"]
            );
        }

        return true;
    }
}

/**
 * @member {String} name
 */
V1EventType.prototype["name"] = undefined;

/**
 * @member {module:model/V1ArtifactKind} kind
 */
V1EventType.prototype["kind"] = undefined;

export default V1EventType;
