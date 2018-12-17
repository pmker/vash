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

function _interopRequireDefault(obj) {
    return obj && obj.__esModule ? obj : { 'default': obj };
}

function _classCallCheck(instance, Constructor) {
    if (!(instance instanceof Constructor)) {
        throw new TypeError('Cannot call a class as a function');
    }
}

var _ApiClient = require('../ApiClient');

var _ApiClient2 = _interopRequireDefault(_ApiClient);

/**
* The InstallInstallResponse model module.
* @module model/InstallInstallResponse
* @version 1.0
*/

var InstallInstallResponse = function () {
    /**
    * Constructs a new <code>InstallInstallResponse</code>.
    * @alias module:model/InstallInstallResponse
    * @class
    */

    function InstallInstallResponse() {
        _classCallCheck(this, InstallInstallResponse);

        this.success = undefined;
    }

    /**
    * Constructs a <code>InstallInstallResponse</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/InstallInstallResponse} obj Optional instance to populate.
    * @return {module:model/InstallInstallResponse} The populated <code>InstallInstallResponse</code> instance.
    */

    InstallInstallResponse.constructFromObject = function constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstallInstallResponse();

            if (data.hasOwnProperty('success')) {
                obj['success'] = _ApiClient2['default'].convertToType(data['success'], 'Boolean');
            }
        }
        return obj;
    };

    /**
    * @member {Boolean} success
    */
    return InstallInstallResponse;
}();

exports['default'] = InstallInstallResponse;
module.exports = exports['default'];
