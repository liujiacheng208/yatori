package com.cbq.brushlessons.core.utils;

import com.cbq.brushlessons.core.entity.Config;
import org.junit.Test;

public class ConfigUtilsTest{

    @Test
    public void load(){
        Config config = ConfigUtils.loadingConfig();
        System.out.println(config.getUsers().size());
//        System.out.println(config.getUsers().get(0).getAccountType());
    }

}