/**
 * @author  cyberjie
 * @date  2020/8/3 16:24
 * @version 1.0
 */
package org.codesample.swagger.controller;

import io.swagger.annotations.*;
import io.swagger.v3.oas.annotations.Parameter;
import io.swagger.v3.oas.annotations.servers.Server;
import org.codesample.swagger.entity.ResultMessage;
import org.codesample.swagger.entity.helper.ResultMessageHelper;
import org.springframework.web.bind.annotation.*;

@RestController
@Api(value = "code",tags = {"code first"})
@Server(url = "$server/code/**",description = "server code region")
@RequestMapping(value = "/code", method = {RequestMethod.POST, RequestMethod.GET})
public class SwaggerController {
    @ApiResponse(code = 200,responseHeaders ={@ResponseHeader(name = "token",description = "token auth")} ,message = "ok",examples = @Example({@ExampleProperty(value ="{'status':'ok','data':'fine'}",mediaType = "/code/first")}))
    @GetMapping("/first")
    public ResultMessage doFirst() {
        return ResultMessageHelper.ok("fine");
    }

    @PostMapping("/second")
    @Parameter(name = "tag",description = "doc tag",required = true)
    public ResultMessage doc(@RequestParam("tag")String tag){
        return ResultMessageHelper.error(tag);
    }
}
