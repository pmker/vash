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


import ApiClient from '../ApiClient';
import NodeChangeEventEventType from './NodeChangeEventEventType';
import TreeNode from './TreeNode';





/**
* The TreeNodeChangeEvent model module.
* @module model/TreeNodeChangeEvent
* @version 1.0
*/
export default class TreeNodeChangeEvent {
    /**
    * Constructs a new <code>TreeNodeChangeEvent</code>.
    * @alias module:model/TreeNodeChangeEvent
    * @class
    */

    constructor() {
        

        
        

        

        
    }

    /**
    * Constructs a <code>TreeNodeChangeEvent</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/TreeNodeChangeEvent} obj Optional instance to populate.
    * @return {module:model/TreeNodeChangeEvent} The populated <code>TreeNodeChangeEvent</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new TreeNodeChangeEvent();

            
            
            

            if (data.hasOwnProperty('Type')) {
                obj['Type'] = NodeChangeEventEventType.constructFromObject(data['Type']);
            }
            if (data.hasOwnProperty('Source')) {
                obj['Source'] = TreeNode.constructFromObject(data['Source']);
            }
            if (data.hasOwnProperty('Target')) {
                obj['Target'] = TreeNode.constructFromObject(data['Target']);
            }
        }
        return obj;
    }

    /**
    * @member {module:model/NodeChangeEventEventType} Type
    */
    Type = undefined;
    /**
    * @member {module:model/TreeNode} Source
    */
    Source = undefined;
    /**
    * @member {module:model/TreeNode} Target
    */
    Target = undefined;








}


