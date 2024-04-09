/**
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.1.6-rc0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.PolyaxonSdk);
  }
}(this, function(expect, PolyaxonSdk) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new PolyaxonSdk.V1RunSettings();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('V1RunSettings', function() {
    it('should create an instance of V1RunSettings', function() {
      // uncomment below and update the code to test V1RunSettings
      //var instance = new PolyaxonSdk.V1RunSettings();
      //expect(instance).to.be.a(PolyaxonSdk.V1RunSettings);
    });

    it('should have the property namespace (base name: "namespace")', function() {
      // uncomment below and update the code to test the property namespace
      //var instance = new PolyaxonSdk.V1RunSettings();
      //expect(instance).to.be();
    });

    it('should have the property agent (base name: "agent")', function() {
      // uncomment below and update the code to test the property agent
      //var instance = new PolyaxonSdk.V1RunSettings();
      //expect(instance).to.be();
    });

    it('should have the property queue (base name: "queue")', function() {
      // uncomment below and update the code to test the property queue
      //var instance = new PolyaxonSdk.V1RunSettings();
      //expect(instance).to.be();
    });

    it('should have the property artifacts_store (base name: "artifacts_store")', function() {
      // uncomment below and update the code to test the property artifacts_store
      //var instance = new PolyaxonSdk.V1RunSettings();
      //expect(instance).to.be();
    });

    it('should have the property tensorboard (base name: "tensorboard")', function() {
      // uncomment below and update the code to test the property tensorboard
      //var instance = new PolyaxonSdk.V1RunSettings();
      //expect(instance).to.be();
    });

    it('should have the property build (base name: "build")', function() {
      // uncomment below and update the code to test the property build
      //var instance = new PolyaxonSdk.V1RunSettings();
      //expect(instance).to.be();
    });

    it('should have the property component (base name: "component")', function() {
      // uncomment below and update the code to test the property component
      //var instance = new PolyaxonSdk.V1RunSettings();
      //expect(instance).to.be();
    });

    it('should have the property models (base name: "models")', function() {
      // uncomment below and update the code to test the property models
      //var instance = new PolyaxonSdk.V1RunSettings();
      //expect(instance).to.be();
    });

    it('should have the property artifacts (base name: "artifacts")', function() {
      // uncomment below and update the code to test the property artifacts
      //var instance = new PolyaxonSdk.V1RunSettings();
      //expect(instance).to.be();
    });

  });

}));
