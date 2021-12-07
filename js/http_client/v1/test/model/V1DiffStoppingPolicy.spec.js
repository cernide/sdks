// Copyright 2018-2021 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.13.0-rc3
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
    instance = new PolyaxonSdk.V1DiffStoppingPolicy();
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

  describe('V1DiffStoppingPolicy', function() {
    it('should create an instance of V1DiffStoppingPolicy', function() {
      // uncomment below and update the code to test V1DiffStoppingPolicy
      //var instance = new PolyaxonSdk.V1DiffStoppingPolicy();
      //expect(instance).to.be.a(PolyaxonSdk.V1DiffStoppingPolicy);
    });

    it('should have the property kind (base name: "kind")', function() {
      // uncomment below and update the code to test the property kind
      //var instance = new PolyaxonSdk.V1DiffStoppingPolicy();
      //expect(instance).to.be();
    });

    it('should have the property percent (base name: "percent")', function() {
      // uncomment below and update the code to test the property percent
      //var instance = new PolyaxonSdk.V1DiffStoppingPolicy();
      //expect(instance).to.be();
    });

    it('should have the property evaluationInterval (base name: "evaluationInterval")', function() {
      // uncomment below and update the code to test the property evaluationInterval
      //var instance = new PolyaxonSdk.V1DiffStoppingPolicy();
      //expect(instance).to.be();
    });

    it('should have the property minInterval (base name: "minInterval")', function() {
      // uncomment below and update the code to test the property minInterval
      //var instance = new PolyaxonSdk.V1DiffStoppingPolicy();
      //expect(instance).to.be();
    });

    it('should have the property minSamples (base name: "minSamples")', function() {
      // uncomment below and update the code to test the property minSamples
      //var instance = new PolyaxonSdk.V1DiffStoppingPolicy();
      //expect(instance).to.be();
    });

  });

}));