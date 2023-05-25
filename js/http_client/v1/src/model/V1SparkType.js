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
* Enum class V1SparkType.
* @enum {}
* @readonly
*/
export default class V1SparkType {
    
        /**
         * value: "java"
         * @const
         */
        "java" = "java";

    
        /**
         * value: "scala"
         * @const
         */
        "scala" = "scala";

    
        /**
         * value: "python"
         * @const
         */
        "python" = "python";

    
        /**
         * value: "r"
         * @const
         */
        "r" = "r";

    

    /**
    * Returns a <code>V1SparkType</code> enum value from a Javascript object name.
    * @param {Object} data The plain JavaScript object containing the name of the enum value.
    * @return {module:model/V1SparkType} The enum <code>V1SparkType</code> value.
    */
    static constructFromObject(object) {
        return object;
    }
}

