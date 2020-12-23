create table recharge
(
    uid           int          null,
    openid        varchar(100) not null,
    payno         varchar(100) not null
        primary key,
    recharge_time varchar(20)  null,
    rmb           varchar(20)  null,
    item_id       int          null,
    channel       varchar(10)  null
);

create table shopmall
(
    id                int         not null
        primary key,
    sellingway        int         null,
    recommend         int         null,
    recommendactivity int         null,
    price             int         null,
    discountprice     double      null,
    starttime         varchar(20) null,
    endtime           varchar(20) null
);

create table userinfo
(
    uid      int                     null,
    openid   varchar(100) default '' not null
        primary key,
    psw      varchar(15)  default '' null,
    mac      varchar(50)  default '' null,
    nickname varchar(30)  default '' null
);

create table useritem
(
    uid       int           null,
    openid    varchar(100)  not null
        primary key,
    shop_list varchar(5000) null
);


