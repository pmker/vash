/**
 * Pydio Cells Rest API
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *
 */

'use strict';

exports.__esModule = true;

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError('Cannot call a class as a function'); } }

var _ApiClient = require("../ApiClient");

var _ApiClient2 = _interopRequireDefault(_ApiClient);

var _modelRestBulkMetaResponse = require('../model/RestBulkMetaResponse');

var _modelRestBulkMetaResponse2 = _interopRequireDefault(_modelRestBulkMetaResponse);

var _modelRestGetBulkMetaRequest = require('../model/RestGetBulkMetaRequest');

var _modelRestGetBulkMetaRequest2 = _interopRequireDefault(_modelRestGetBulkMetaRequest);

var _modelRestMetaCollection = require('../model/RestMetaCollection');

var _modelRestMetaCollection2 = _interopRequireDefault(_modelRestMetaCollection);

var _modelRestMetaNamespaceRequest = require('../model/RestMetaNamespaceRequest');

var _modelRestMetaNamespaceRequest2 = _interopRequireDefault(_modelRestMetaNamespaceRequest);

var _modelTreeNode = require('../model/TreeNode');

var _modelTreeNode2 = _interopRequireDefault(_modelTreeNode);

/**
* MetaService service.
* @module api/MetaServiceApi
* @version 1.0
*/

var MetaServiceApi = (function () {

  /**
  * Constructs a new MetaServiceApi. 
  * @alias module:api/MetaServiceApi
  * @class
  * @param {module:ApiClient} apiClient Optional API client implementation to use,
  * default to {@link module:ApiClient#instance} if unspecified.
  */

  function MetaServiceApi(apiClient) {
    _classCallCheck(this, MetaServiceApi);

    this.apiClient = apiClient || _ApiClient2['default'].instance;
  }

  /**
   * Delete metadata of a given node
   * @param {String} nodePath 
   * @param {module:model/RestMetaNamespaceRequest} body 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/TreeNode} and HTTP response
   */

  MetaServiceApi.prototype.deleteMetaWithHttpInfo = function deleteMetaWithHttpInfo(nodePath, body) {
    var postBody = body;

    // verify the required parameter 'nodePath' is set
    if (nodePath === undefined || nodePath === null) {
      throw new Error("Missing the required parameter 'nodePath' when calling deleteMeta");
    }

    // verify the required parameter 'body' is set
    if (body === undefined || body === null) {
      throw new Error("Missing the required parameter 'body' when calling deleteMeta");
    }

    var pathParams = {
      'NodePath': nodePath
    };
    var queryParams = {};
    var headerParams = {};
    var formParams = {};

    var authNames = [];
    var contentTypes = ['application/json'];
    var accepts = ['application/json'];
    var returnType = _modelTreeNode2['default'];

    return this.apiClient.callApi('/meta/delete/{NodePath}', 'POST', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType);
  };

  /**
   * Delete metadata of a given node
   * @param {String} nodePath 
   * @param {module:model/RestMetaNamespaceRequest} body 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/TreeNode}
   */

  MetaServiceApi.prototype.deleteMeta = function deleteMeta(nodePath, body) {
    return this.deleteMetaWithHttpInfo(nodePath, body).then(function (response_and_data) {
      return response_and_data.data;
    });
  };

  /**
   * List meta for a list of nodes, or a full directory using /path/_* syntax
   * @param {module:model/RestGetBulkMetaRequest} body 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/RestBulkMetaResponse} and HTTP response
   */

  MetaServiceApi.prototype.getBulkMetaWithHttpInfo = function getBulkMetaWithHttpInfo(body) {
    var postBody = body;

    // verify the required parameter 'body' is set
    if (body === undefined || body === null) {
      throw new Error("Missing the required parameter 'body' when calling getBulkMeta");
    }

    var pathParams = {};
    var queryParams = {};
    var headerParams = {};
    var formParams = {};

    var authNames = [];
    var contentTypes = ['application/json'];
    var accepts = ['application/json'];
    var returnType = _modelRestBulkMetaResponse2['default'];

    return this.apiClient.callApi('/meta/bulk/get', 'POST', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType);
  };

  /**
   * List meta for a list of nodes, or a full directory using /path/_* syntax
   * @param {module:model/RestGetBulkMetaRequest} body 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/RestBulkMetaResponse}
   */

  MetaServiceApi.prototype.getBulkMeta = function getBulkMeta(body) {
    return this.getBulkMetaWithHttpInfo(body).then(function (response_and_data) {
      return response_and_data.data;
    });
  };

  /**
   * Load metadata for a given node
   * @param {String} nodePath 
   * @param {module:model/RestMetaNamespaceRequest} body 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/TreeNode} and HTTP response
   */

  MetaServiceApi.prototype.getMetaWithHttpInfo = function getMetaWithHttpInfo(nodePath, body) {
    var postBody = body;

    // verify the required parameter 'nodePath' is set
    if (nodePath === undefined || nodePath === null) {
      throw new Error("Missing the required parameter 'nodePath' when calling getMeta");
    }

    // verify the required parameter 'body' is set
    if (body === undefined || body === null) {
      throw new Error("Missing the required parameter 'body' when calling getMeta");
    }

    var pathParams = {
      'NodePath': nodePath
    };
    var queryParams = {};
    var headerParams = {};
    var formParams = {};

    var authNames = [];
    var contentTypes = ['application/json'];
    var accepts = ['application/json'];
    var returnType = _modelTreeNode2['default'];

    return this.apiClient.callApi('/meta/get/{NodePath}', 'POST', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType);
  };

  /**
   * Load metadata for a given node
   * @param {String} nodePath 
   * @param {module:model/RestMetaNamespaceRequest} body 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/TreeNode}
   */

  MetaServiceApi.prototype.getMeta = function getMeta(nodePath, body) {
    return this.getMetaWithHttpInfo(nodePath, body).then(function (response_and_data) {
      return response_and_data.data;
    });
  };

  /**
   * Update metadata for a given node
   * @param {String} nodePath 
   * @param {module:model/RestMetaCollection} body 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/TreeNode} and HTTP response
   */

  MetaServiceApi.prototype.setMetaWithHttpInfo = function setMetaWithHttpInfo(nodePath, body) {
    var postBody = body;

    // verify the required parameter 'nodePath' is set
    if (nodePath === undefined || nodePath === null) {
      throw new Error("Missing the required parameter 'nodePath' when calling setMeta");
    }

    // verify the required parameter 'body' is set
    if (body === undefined || body === null) {
      throw new Error("Missing the required parameter 'body' when calling setMeta");
    }

    var pathParams = {
      'NodePath': nodePath
    };
    var queryParams = {};
    var headerParams = {};
    var formParams = {};

    var authNames = [];
    var contentTypes = ['application/json'];
    var accepts = ['application/json'];
    var returnType = _modelTreeNode2['default'];

    return this.apiClient.callApi('/meta/set/{NodePath}', 'POST', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType);
  };

  /**
   * Update metadata for a given node
   * @param {String} nodePath 
   * @param {module:model/RestMetaCollection} body 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/TreeNode}
   */

  MetaServiceApi.prototype.setMeta = function setMeta(nodePath, body) {
    return this.setMetaWithHttpInfo(nodePath, body).then(function (response_and_data) {
      return response_and_data.data;
    });
  };

  return MetaServiceApi;
})();

exports['default'] = MetaServiceApi;
module.exports = exports['default'];
