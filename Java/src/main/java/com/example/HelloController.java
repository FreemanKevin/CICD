package com.example;

import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@Api(tags = "Hello接口")
@RestController
public class HelloController {
    
    @ApiOperation("获取问候语")
    @GetMapping("/hello")
    public String hello() {
        return "Hello Swagger UI!";
    }
} 