/**
 * @author  cyberjie
 * @date  2020/8/3 16:25
 * @version 1.0
 */
package org.codesample.swagger.entity.helper;

import org.codesample.swagger.entity.ResultMessage;

public class ResultMessageHelper {
    private static final String STATUS_OK = "ok";
    private static final String STATUS_ERROR = "error";

    public static ResultMessage ok(Object data) {
        return new ResultMessage(STATUS_OK, data);
    }

    public static ResultMessage error(Object data) {
        return new ResultMessage(STATUS_ERROR, data);
    }
}
