import time

import requests
from bs4 import BeautifulSoup as bs

basic_url = 'http://category.dangdang.com'


# 每个类目一个爬取器
class Crawler:
    def __init__(self, name, url, num) -> None:
        self.name = name
        self.url = url
        self.num = num
        self.data = []
        self.cid = 0
        self.cid2 = 0


    def run(self):
        url = self.url
        while len(self.data) < self.num:
            list, next_url = spider_once(url)
            # print('list len={} cur_len={}'.format(len(list), len(self.data)), flush=True)
            if list is None or len(list) == 0:
                continue
            self.data.extend(list)
            if next_url != "":
                url = next_url
            time.sleep(0.5)


def spider_once(url):
    """ 请求一次，解析html """
    print(url, flush=True)
    r = requests.get(url=url, timeout=5)
    # 使用 lxml 解析器
    soup = bs(r.text, 'lxml')
    list = soup.find_all('li')
    list = soup.select('ul.bigimg > li')
    li = list[0]

    resp = []

    for li in list:
        book = parse_book_info(li)
        if book is None:
            continue
        resp.append(book)

    # 下一页的url
    next = soup.select_one('ul[name=Fy] > .next').a.get('href')
    next_url = basic_url + next

    return resp, next_url


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
    authorItem = parse_author(author_publish_tag)
    if authorItem is None:
        print('parse book error: ', title, author_publish_tag, flush=True)
        return None

    return {
        'title': title,
        'author': authorItem['author'],
        'publisher': authorItem['publisher'],
        'publish_time': authorItem['publish_time'],
        'image': image,
        'price': float(pre_price),
        'cur_price': float(cur_price),
        'info_url': info_url,
        'detail': detail,
    }


def parse_author(tag) -> dict:
    spans = tag.find_all('span')
    if len(spans) != 3:
        print('--------- author_publish_tag error --------', flush=True)
        return None

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
        'author': author,
        'publisher': publisher,
        'publish_time': publish_time,
    }
