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





/**
* The RestUserJobResponse model module.
* @module model/RestUserJobResponse
* @version 1.0
*/
export default class RestUserJobResponse {
    /**
    * Constructs a new <code>RestUserJobResponse</code>.
    * @alias module:model/RestUserJobResponse
    * @class
    */

    constructor() {
        

        
        

        

        
    }

    /**
    * Constructs a <code>RestUserJobResponse</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/RestUserJobResponse} obj Optional instance to populate.
    * @return {module:model/RestUserJobResponse} The populated <code>RestUserJobResponse</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RestUserJobResponse();

            
            
            

            if (data.hasOwnProperty('JobUuid')) {
                obj['JobUuid'] = ApiClient.convertToType(data['JobUuid'], 'String');
            }
        }
        return obj;
    }

    /**
    * @member {String} JobUuid
    */
    JobUuid = undefined;








}


