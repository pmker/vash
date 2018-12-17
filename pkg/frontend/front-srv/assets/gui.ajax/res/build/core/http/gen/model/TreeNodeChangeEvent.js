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

var _NodeChangeEventEventType = require('./NodeChangeEventEventType');

var _NodeChangeEventEventType2 = _interopRequireDefault(_NodeChangeEventEventType);

var _TreeNode = require('./TreeNode');

var _TreeNode2 = _interopRequireDefault(_TreeNode);

/**
* The TreeNodeChangeEvent model module.
* @module model/TreeNodeChangeEvent
* @version 1.0
*/

var TreeNodeChangeEvent = (function () {
    /**
    * Constructs a new <code>TreeNodeChangeEvent</code>.
    * @alias module:model/TreeNodeChangeEvent
    * @class
    */

    function TreeNodeChangeEvent() {
        _classCallCheck(this, TreeNodeChangeEvent);

        this.Type = undefined;
        this.Source = undefined;
        this.Target = undefined;
    }

    /**
    * Constructs a <code>TreeNodeChangeEvent</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/TreeNodeChangeEvent} obj Optional instance to populate.
    * @return {module:model/TreeNodeChangeEvent} The populated <code>TreeNodeChangeEvent</code> instance.
    */

    TreeNodeChangeEvent.constructFromObject = function constructFromObject(data, obj) {
        if (data) {
            obj = obj || new TreeNodeChangeEvent();

            if (data.hasOwnProperty('Type')) {
                obj['Type'] = _NodeChangeEventEventType2['default'].constructFromObject(data['Type']);
            }
            if (data.hasOwnProperty('Source')) {
                obj['Source'] = _TreeNode2['default'].constructFromObject(data['Source']);
            }
            if (data.hasOwnProperty('Target')) {
                obj['Target'] = _TreeNode2['default'].constructFromObject(data['Target']);
            }
        }
        return obj;
    };

    /**
    * @member {module:model/NodeChangeEventEventType} Type
    */
    return TreeNodeChangeEvent;
})();

exports['default'] = TreeNodeChangeEvent;
module.exports = exports['default'];

/**
* @member {module:model/TreeNode} Source
*/

/**
* @member {module:model/TreeNode} Target
*/
