/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc19
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
    instance = new PolyaxonSdk.V1SearchSpec();
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

  describe('V1SearchSpec', function() {
    it('should create an instance of V1SearchSpec', function() {
      // uncomment below and update the code to test V1SearchSpec
      //var instance = new PolyaxonSdk.V1SearchSpec();
      //expect(instance).to.be.a(PolyaxonSdk.V1SearchSpec);
    });

    it('should have the property query (base name: "query")', function() {
      // uncomment below and update the code to test the property query
      //var instance = new PolyaxonSdk.V1SearchSpec();
      //expect(instance).to.be();
    });

    it('should have the property sort (base name: "sort")', function() {
      // uncomment below and update the code to test the property sort
      //var instance = new PolyaxonSdk.V1SearchSpec();
      //expect(instance).to.be();
    });

    it('should have the property limit (base name: "limit")', function() {
      // uncomment below and update the code to test the property limit
      //var instance = new PolyaxonSdk.V1SearchSpec();
      //expect(instance).to.be();
    });

    it('should have the property offset (base name: "offset")', function() {
      // uncomment below and update the code to test the property offset
      //var instance = new PolyaxonSdk.V1SearchSpec();
      //expect(instance).to.be();
    });

    it('should have the property groupby (base name: "groupby")', function() {
      // uncomment below and update the code to test the property groupby
      //var instance = new PolyaxonSdk.V1SearchSpec();
      //expect(instance).to.be();
    });

    it('should have the property columns (base name: "columns")', function() {
      // uncomment below and update the code to test the property columns
      //var instance = new PolyaxonSdk.V1SearchSpec();
      //expect(instance).to.be();
    });

    it('should have the property layout (base name: "layout")', function() {
      // uncomment below and update the code to test the property layout
      //var instance = new PolyaxonSdk.V1SearchSpec();
      //expect(instance).to.be();
    });

    it('should have the property sections (base name: "sections")', function() {
      // uncomment below and update the code to test the property sections
      //var instance = new PolyaxonSdk.V1SearchSpec();
      //expect(instance).to.be();
    });

    it('should have the property compares (base name: "compares")', function() {
      // uncomment below and update the code to test the property compares
      //var instance = new PolyaxonSdk.V1SearchSpec();
      //expect(instance).to.be();
    });

    it('should have the property heat (base name: "heat")', function() {
      // uncomment below and update the code to test the property heat
      //var instance = new PolyaxonSdk.V1SearchSpec();
      //expect(instance).to.be();
    });

    it('should have the property events (base name: "events")', function() {
      // uncomment below and update the code to test the property events
      //var instance = new PolyaxonSdk.V1SearchSpec();
      //expect(instance).to.be();
    });

    it('should have the property histograms (base name: "histograms")', function() {
      // uncomment below and update the code to test the property histograms
      //var instance = new PolyaxonSdk.V1SearchSpec();
      //expect(instance).to.be();
    });

    it('should have the property trends (base name: "trends")', function() {
      // uncomment below and update the code to test the property trends
      //var instance = new PolyaxonSdk.V1SearchSpec();
      //expect(instance).to.be();
    });

    it('should have the property analytics (base name: "analytics")', function() {
      // uncomment below and update the code to test the property analytics
      //var instance = new PolyaxonSdk.V1SearchSpec();
      //expect(instance).to.be();
    });

  });

}));
