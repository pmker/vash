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

var _ApiClient = require('../ApiClient');

var _ApiClient2 = _interopRequireDefault(_ApiClient);

var _TreeVersioningKeepPeriod = require('./TreeVersioningKeepPeriod');

var _TreeVersioningKeepPeriod2 = _interopRequireDefault(_TreeVersioningKeepPeriod);

/**
* The TreeVersioningPolicy model module.
* @module model/TreeVersioningPolicy
* @version 1.0
*/

var TreeVersioningPolicy = (function () {
    /**
    * Constructs a new <code>TreeVersioningPolicy</code>.
    * @alias module:model/TreeVersioningPolicy
    * @class
    */

    function TreeVersioningPolicy() {
        _classCallCheck(this, TreeVersioningPolicy);

        this.Uuid = undefined;
        this.Name = undefined;
        this.Description = undefined;
        this.VersionsDataSourceName = undefined;
        this.VersionsDataSourceBucket = undefined;
        this.MaxTotalSize = undefined;
        this.MaxSizePerFile = undefined;
        this.IgnoreFilesGreaterThan = undefined;
        this.KeepPeriods = undefined;
    }

    /**
    * Constructs a <code>TreeVersioningPolicy</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/TreeVersioningPolicy} obj Optional instance to populate.
    * @return {module:model/TreeVersioningPolicy} The populated <code>TreeVersioningPolicy</code> instance.
    */

    TreeVersioningPolicy.constructFromObject = function constructFromObject(data, obj) {
        if (data) {
            obj = obj || new TreeVersioningPolicy();

            if (data.hasOwnProperty('Uuid')) {
                obj['Uuid'] = _ApiClient2['default'].convertToType(data['Uuid'], 'String');
            }
            if (data.hasOwnProperty('Name')) {
                obj['Name'] = _ApiClient2['default'].convertToType(data['Name'], 'String');
            }
            if (data.hasOwnProperty('Description')) {
                obj['Description'] = _ApiClient2['default'].convertToType(data['Description'], 'String');
            }
            if (data.hasOwnProperty('VersionsDataSourceName')) {
                obj['VersionsDataSourceName'] = _ApiClient2['default'].convertToType(data['VersionsDataSourceName'], 'String');
            }
            if (data.hasOwnProperty('VersionsDataSourceBucket')) {
                obj['VersionsDataSourceBucket'] = _ApiClient2['default'].convertToType(data['VersionsDataSourceBucket'], 'String');
            }
            if (data.hasOwnProperty('MaxTotalSize')) {
                obj['MaxTotalSize'] = _ApiClient2['default'].convertToType(data['MaxTotalSize'], 'String');
            }
            if (data.hasOwnProperty('MaxSizePerFile')) {
                obj['MaxSizePerFile'] = _ApiClient2['default'].convertToType(data['MaxSizePerFile'], 'String');
            }
            if (data.hasOwnProperty('IgnoreFilesGreaterThan')) {
                obj['IgnoreFilesGreaterThan'] = _ApiClient2['default'].convertToType(data['IgnoreFilesGreaterThan'], 'String');
            }
            if (data.hasOwnProperty('KeepPeriods')) {
                obj['KeepPeriods'] = _ApiClient2['default'].convertToType(data['KeepPeriods'], [_TreeVersioningKeepPeriod2['default']]);
            }
        }
        return obj;
    };

    /**
    * @member {String} Uuid
    */
    return TreeVersioningPolicy;
})();

exports['default'] = TreeVersioningPolicy;
module.exports = exports['default'];

/**
* @member {String} Name
*/

/**
* @member {String} Description
*/

/**
* @member {String} VersionsDataSourceName
*/

/**
* @member {String} VersionsDataSourceBucket
*/

/**
* @member {String} MaxTotalSize
*/

/**
* @member {String} MaxSizePerFile
*/

/**
* @member {String} IgnoreFilesGreaterThan
*/

/**
* @member {Array.<module:model/TreeVersioningKeepPeriod>} KeepPeriods
*/
