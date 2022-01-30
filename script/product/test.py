# -*- coding: utf-8 -*-

import io
import sys

import requests
from bs4 import BeautifulSoup as bs

# print 乱码解决方法
# 参考：http://www.360doc.com/content/19/1018/19/12906439_867676902.shtml
sys.stdout = io.TextIOWrapper(sys.stdout.buffer, encoding='utf-8')

basic_url = 'http://category.dangdang.com'

def run():
    url = basic_url + '/cp01.01.00.00.00.00.html'
    r = requests.get(url=url)
    # 使用 lxml 解析器
    soup = bs(r.text, 'lxml')
    list = soup.find_all('li')
    list = soup.select('ul.bigimg > li')
    sum = 0
    li = list[0]
    next = soup.select_one('ul[name=Fy] > .next').a.get('href')
    print(next)

    for li in list:
        sum += 1
        book = parse_book_info(li)
        print(book)

    print(sum)

def parse_book_info(li) -> dict:
    # 标题
    title = li.a.get('title').strip()
    # 图片url
    image = "http:" + li.a.img.get('src')
    #  图书详情地址
    info_url = "http:" + li.a.get('href')
    # 详情介绍
    # todo：去除尾部的“推荐您购买……”
    detail_tag = li.select_one('p.detail')
    detail = detail_tag.text.replace('★', '').replace('◆', '').replace('�h', '')

    # 价格
    price_tag = li.select_one('p.price')
    cur_price = price_tag.select_one('span.search_now_price').text.strip().lstrip('¥')
    pre_price = price_tag.select_one('span.search_pre_price').text.strip().lstrip('¥')

    # 作者、出版社和出版时间
    author_publish_tag = li.select_one('p.search_book_author')
    spans = author_publish_tag.find_all('span')
    if len(spans) != 3:
        print('--------- author_publish_tag error --------')
        return {}

    # 有三个 span，如：景行 白马时光 出品 /2020-04-01  /百花文艺出版社
    # 第一个span包含作者、译者、出品组织，可能内容以<a>标签包含，也可能无
    # 第二个span为出版时间
    # 作者
    author_tag = spans[0].find('a')
    if author_tag is None:
        author = spans[0].text
    else:
        author = author_tag.text
    # 出版社
    publisher = spans[2].text.strip().lstrip('/')
    # 出版时间
    publish_time = spans[1].text.strip().lstrip('/')

    return {
        'title': title,
        'author': author,
        'publisher': publisher,
        'publish_time': publish_time,
        'image': image,
        'cur_price': cur_price,
        'pre_price': pre_price,
        'info_url': info_url,
        'detail': detail,
    }


if __name__ == '__main__':
    print(sys.getdefaultencoding())
    print('run')
    run()
