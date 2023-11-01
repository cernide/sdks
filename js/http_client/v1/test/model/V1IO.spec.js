/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0
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
    instance = new PolyaxonSdk.V1IO();
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

  describe('V1IO', function() {
    it('should create an instance of V1IO', function() {
      // uncomment below and update the code to test V1IO
      //var instance = new PolyaxonSdk.V1IO();
      //expect(instance).to.be.a(PolyaxonSdk.V1IO);
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new PolyaxonSdk.V1IO();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new PolyaxonSdk.V1IO();
      //expect(instance).to.be();
    });

    it('should have the property type (base name: "type")', function() {
      // uncomment below and update the code to test the property type
      //var instance = new PolyaxonSdk.V1IO();
      //expect(instance).to.be();
    });

    it('should have the property value (base name: "value")', function() {
      // uncomment below and update the code to test the property value
      //var instance = new PolyaxonSdk.V1IO();
      //expect(instance).to.be();
    });

    it('should have the property isOptional (base name: "isOptional")', function() {
      // uncomment below and update the code to test the property isOptional
      //var instance = new PolyaxonSdk.V1IO();
      //expect(instance).to.be();
    });

    it('should have the property isList (base name: "isList")', function() {
      // uncomment below and update the code to test the property isList
      //var instance = new PolyaxonSdk.V1IO();
      //expect(instance).to.be();
    });

    it('should have the property isFlag (base name: "isFlag")', function() {
      // uncomment below and update the code to test the property isFlag
      //var instance = new PolyaxonSdk.V1IO();
      //expect(instance).to.be();
    });

    it('should have the property argFormat (base name: "argFormat")', function() {
      // uncomment below and update the code to test the property argFormat
      //var instance = new PolyaxonSdk.V1IO();
      //expect(instance).to.be();
    });

    it('should have the property connection (base name: "connection")', function() {
      // uncomment below and update the code to test the property connection
      //var instance = new PolyaxonSdk.V1IO();
      //expect(instance).to.be();
    });

    it('should have the property toInit (base name: "toInit")', function() {
      // uncomment below and update the code to test the property toInit
      //var instance = new PolyaxonSdk.V1IO();
      //expect(instance).to.be();
    });

    it('should have the property toEnv (base name: "toEnv")', function() {
      // uncomment below and update the code to test the property toEnv
      //var instance = new PolyaxonSdk.V1IO();
      //expect(instance).to.be();
    });

    it('should have the property validation (base name: "validation")', function() {
      // uncomment below and update the code to test the property validation
      //var instance = new PolyaxonSdk.V1IO();
      //expect(instance).to.be();
    });

    it('should have the property delayValidation (base name: "delayValidation")', function() {
      // uncomment below and update the code to test the property delayValidation
      //var instance = new PolyaxonSdk.V1IO();
      //expect(instance).to.be();
    });

    it('should have the property options (base name: "options")', function() {
      // uncomment below and update the code to test the property options
      //var instance = new PolyaxonSdk.V1IO();
      //expect(instance).to.be();
    });

  });

}));
