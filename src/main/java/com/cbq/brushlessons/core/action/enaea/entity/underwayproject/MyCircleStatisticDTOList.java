package com.cbq.brushlessons.core.action.enaea.entity.underwayproject;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

@lombok.Data
public class MyCircleStatisticDTOList {
    @lombok.Getter(onMethod_ = {@JsonProperty("list")})
    @lombok.Setter(onMethod_ = {@JsonProperty("list")})
    private List<String> list;
}