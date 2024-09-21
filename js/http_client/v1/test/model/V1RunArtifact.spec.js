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
        instance = new PolyaxonSdk.V1RunArtifact();
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

    describe("V1RunArtifact", function () {
        it("should create an instance of V1RunArtifact", function () {
            // uncomment below and update the code to test V1RunArtifact
            //var instance = new PolyaxonSdk.V1RunArtifact();
            //expect(instance).to.be.a(PolyaxonSdk.V1RunArtifact);
        });

        it('should have the property name (base name: "name")', function () {
            // uncomment below and update the code to test the property name
            //var instance = new PolyaxonSdk.V1RunArtifact();
            //expect(instance).to.be();
        });

        it('should have the property state (base name: "state")', function () {
            // uncomment below and update the code to test the property state
            //var instance = new PolyaxonSdk.V1RunArtifact();
            //expect(instance).to.be();
        });

        it('should have the property kind (base name: "kind")', function () {
            // uncomment below and update the code to test the property kind
            //var instance = new PolyaxonSdk.V1RunArtifact();
            //expect(instance).to.be();
        });

        it('should have the property path (base name: "path")', function () {
            // uncomment below and update the code to test the property path
            //var instance = new PolyaxonSdk.V1RunArtifact();
            //expect(instance).to.be();
        });

        it('should have the property connection (base name: "connection")', function () {
            // uncomment below and update the code to test the property connection
            //var instance = new PolyaxonSdk.V1RunArtifact();
            //expect(instance).to.be();
        });

        it('should have the property run (base name: "run")', function () {
            // uncomment below and update the code to test the property run
            //var instance = new PolyaxonSdk.V1RunArtifact();
            //expect(instance).to.be();
        });

        it('should have the property summary (base name: "summary")', function () {
            // uncomment below and update the code to test the property summary
            //var instance = new PolyaxonSdk.V1RunArtifact();
            //expect(instance).to.be();
        });

        it('should have the property is_input (base name: "is_input")', function () {
            // uncomment below and update the code to test the property is_input
            //var instance = new PolyaxonSdk.V1RunArtifact();
            //expect(instance).to.be();
        });

        it('should have the property meta_info (base name: "meta_info")', function () {
            // uncomment below and update the code to test the property meta_info
            //var instance = new PolyaxonSdk.V1RunArtifact();
            //expect(instance).to.be();
        });
    });
});
