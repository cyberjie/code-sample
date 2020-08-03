/**
 * @author  cyberjie
 * @date  2020/8/3 16:24
 * @version 1.0
 */
package org.codesample.swagger.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import springfox.documentation.builders.ApiInfoBuilder;
import springfox.documentation.builders.PathSelectors;
import springfox.documentation.builders.RequestHandlerSelectors;
import springfox.documentation.service.Contact;
import springfox.documentation.spi.DocumentationType;
import springfox.documentation.spring.web.plugins.Docket;
import springfox.documentation.swagger2.annotations.EnableSwagger2;

@EnableSwagger2
@Configuration
public class SwaggerConfiguration {
    @Bean
    public Docket getDocket() {
        return new Docket(DocumentationType.SWAGGER_2)
                .apiInfo(new ApiInfoBuilder()
                        .title("code_sample swagger")
                        .description("code_sample swagger api")
                        .license("MIT")
                        .version("1.0")
                        .contact(new Contact("cyberjie", "https://github.com/cyberjie", "cyberjie@163.com"))
                        .build())
                .select()
                .apis(RequestHandlerSelectors.basePackage("org.codesample.swagger"))
                .paths(PathSelectors.any())
                .build();
    }
}
