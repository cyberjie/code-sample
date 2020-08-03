/**
 * @author  cyberjie
 * @date  2020/8/3 16:24
 * @version 1.0
 */
package org.codesample.swagger.entity;

public class ResultMessage {
    String status;
    Object data;

    public ResultMessage(String status, Object data) {
        this.status = status;
        this.data = data;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public Object getData() {
        return data;
    }

    public void setData(Object data) {
        this.data = data;
    }
}
