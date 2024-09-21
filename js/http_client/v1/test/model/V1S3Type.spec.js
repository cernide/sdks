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

(function (root, factory) {
    if (typeof define === "function" && define.amd) {
        // AMD.
        define(["expect.js", process.cwd() + "/src/index"], factory);
    } else if (typeof module === "object" && module.exports) {
        // CommonJS-like environments that support module.exports, like Node.
        factory(require("expect.js"), require(process.cwd() + "/src/index"));
    } else {
        // Browser globals (root is window)
        factory(root.expect, root.PolyaxonSdk);
    }
})(this, function (expect, PolyaxonSdk) {
    "use strict";

    var instance;

    beforeEach(function () {
        instance = new PolyaxonSdk.V1S3Type();
    });

    var getProperty = function (object, getter, property) {
        // Use getter method if present; otherwise, get the property directly.
        if (typeof object[getter] === "function") return object[getter]();
        else return object[property];
    };

    var setProperty = function (object, setter, property, value) {
        // Use setter method if present; otherwise, set the property directly.
        if (typeof object[setter] === "function") object[setter](value);
        else object[property] = value;
    };

    describe("V1S3Type", function () {
        it("should create an instance of V1S3Type", function () {
            // uncomment below and update the code to test V1S3Type
            //var instance = new PolyaxonSdk.V1S3Type();
            //expect(instance).to.be.a(PolyaxonSdk.V1S3Type);
        });

        it('should have the property bucket (base name: "bucket")', function () {
            // uncomment below and update the code to test the property bucket
            //var instance = new PolyaxonSdk.V1S3Type();
            //expect(instance).to.be();
        });

        it('should have the property key (base name: "key")', function () {
            // uncomment below and update the code to test the property key
            //var instance = new PolyaxonSdk.V1S3Type();
            //expect(instance).to.be();
        });
    });
});
