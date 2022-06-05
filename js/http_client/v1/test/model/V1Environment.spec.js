// Copyright 2018-2022 Polyaxon, Inc.
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
 * The version of the OpenAPI document: 1.18.2
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
    instance = new PolyaxonSdk.V1Environment();
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

  describe('V1Environment', function() {
    it('should create an instance of V1Environment', function() {
      // uncomment below and update the code to test V1Environment
      //var instance = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be.a(PolyaxonSdk.V1Environment);
    });

    it('should have the property labels (base name: "labels")', function() {
      // uncomment below and update the code to test the property labels
      //var instance = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property annotations (base name: "annotations")', function() {
      // uncomment below and update the code to test the property annotations
      //var instance = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property nodeSelector (base name: "nodeSelector")', function() {
      // uncomment below and update the code to test the property nodeSelector
      //var instance = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property affinity (base name: "affinity")', function() {
      // uncomment below and update the code to test the property affinity
      //var instance = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property tolerations (base name: "tolerations")', function() {
      // uncomment below and update the code to test the property tolerations
      //var instance = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property nodeName (base name: "nodeName")', function() {
      // uncomment below and update the code to test the property nodeName
      //var instance = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property serviceAccountName (base name: "serviceAccountName")', function() {
      // uncomment below and update the code to test the property serviceAccountName
      //var instance = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property hostAliases (base name: "hostAliases")', function() {
      // uncomment below and update the code to test the property hostAliases
      //var instance = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property securityContext (base name: "securityContext")', function() {
      // uncomment below and update the code to test the property securityContext
      //var instance = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property imagePullSecrets (base name: "imagePullSecrets")', function() {
      // uncomment below and update the code to test the property imagePullSecrets
      //var instance = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property hostNetwork (base name: "hostNetwork")', function() {
      // uncomment below and update the code to test the property hostNetwork
      //var instance = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property hostPID (base name: "hostPID")', function() {
      // uncomment below and update the code to test the property hostPID
      //var instance = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property dnsPolicy (base name: "dnsPolicy")', function() {
      // uncomment below and update the code to test the property dnsPolicy
      //var instance = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property dnsConfig (base name: "dnsConfig")', function() {
      // uncomment below and update the code to test the property dnsConfig
      //var instance = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property schedulerName (base name: "schedulerName")', function() {
      // uncomment below and update the code to test the property schedulerName
      //var instance = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property priorityClassName (base name: "priorityClassName")', function() {
      // uncomment below and update the code to test the property priorityClassName
      //var instance = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property priority (base name: "priority")', function() {
      // uncomment below and update the code to test the property priority
      //var instance = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

    it('should have the property restartPolicy (base name: "restartPolicy")', function() {
      // uncomment below and update the code to test the property restartPolicy
      //var instance = new PolyaxonSdk.V1Environment();
      //expect(instance).to.be();
    });

  });

}));
