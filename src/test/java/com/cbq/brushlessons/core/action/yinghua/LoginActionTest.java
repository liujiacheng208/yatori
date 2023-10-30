package com.cbq.brushlessons.core.action.yinghua;

import com.cbq.brushlessons.core.action.yinghua.entity.allcourse.CourseRequest;
import com.cbq.brushlessons.core.entity.AccountCacheYingHua;
import com.cbq.brushlessons.core.entity.AccountType;
import com.cbq.brushlessons.core.entity.User;
import com.cbq.brushlessons.core.utils.FileUtils;
import com.cbq.brushlessons.core.utils.VerificationCodeUtil;
import org.junit.Test;

import java.io.File;

public class LoginActionTest{

    @Test
    public void getCode(){
        File code = LoginAction.getCode(new User());
        System.out.println(code.getName());
        String s = VerificationCodeUtil.aiDiscern(code);
        FileUtils.deleteFile(code);
    }
    @Test
    public void getSESSION(){
    }

    /**
     * 测试整套登录流程
     */
    @Test
    public void testLogin(){
        //构建用户信息
        User user = new User();
        user.setAccountType(AccountType.YINGHUA);
        user.setUrl("https://mooc.bwgl.cn/");
        user.setAccount("2151118");
        user.setPassword("02Y4Qtvk");

        AccountCacheYingHua accountYingHua = new AccountCacheYingHua();

        user.setCache(accountYingHua);



        //获取SESSION
        String session = LoginAction.getSESSION(user);
        accountYingHua.setSession(session);

        //获取验证码
        File code = LoginAction.getCode(user);
        String s = VerificationCodeUtil.aiDiscern(code);
        accountYingHua.setCode(s);
        System.out.println(s);

        //进行登录操作
        LoginAction.toLogin(user);
        System.out.println(session);
        accountYingHua.setToken("sid.yTcILDi1VwYD1sdfVmvQXIHY7hM1zu");

        //获取全部的课表请求
        CourseRequest allCourseList = CourseAction.getAllCourseRequest(user);

        //将第一门课加入学习
        CourseStudyAction build = CourseStudyAction.builder()
                .user(user)
                .courseInform(allCourseList.getResult().getList().get(6))
                .build();
        build.toStudy();

    }
}