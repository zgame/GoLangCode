-----------------------------------------------
--- 天气
--    -1    全天气    All
--    1	晴天	SunShine
--    2	雨天	Rain
--    3	沙尘暴	SandStorm
--    4	阴天	OvercastSky
--    5	炎热天	HotDay
-----------------------------------------------


function SandRockRoom:GetWeather()
    return self.weather
end

function SandRockRoom:UpdateWeather()
    -- 目前都是晴天， 任务触发下雨和沙尘暴
    self.weather = ZRandom.GetRandom(1,5)
    return self.weather
end