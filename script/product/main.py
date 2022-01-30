import csv
import io
import random
import sys
import time
from tokenize import Number

from crawler import Crawler
from data import target_list

sys.stdout = io.TextIOWrapper(sys.stdout.buffer, encoding='utf-8')

class Service:
    def __init__(self, target_list) -> None:
        self.target_list = target_list
        self.data = []
        # 数据标题集合，用来去重
        self.data_set = set()

    def run(self):
        # 对每个类目进行爬取
        for target in self.target_list:
            name, url, num = target['name'], target['url'], target['num']
            print('-------------- {} is crawling: {}, {}'.format(name, url, num), flush=True)
            crawler = Crawler(name, url, num)
            crawler.run()

            cid = target['cid']
            cid2 = target['cid2']

            # 添加元素
            cur_list = crawler.data
            for item in cur_list:
                key = item['author'] + item['title']
                if key in self.data_set:
                    continue
                self.data_set.add(key)
                self.data.append({
                    'cid': cid,
                    'cid2': cid2,
                    'title': item['title'],
                    'author': item['author'],
                    'publisher': item['publisher'],
                    'publish_time': item['publish_time'],
                    'image': item['image'],
                    'price': item['price'],
                    'cur_price': item['cur_price'],
                    'info_url': item['info_url'],
                    'detail': item['detail'],
                })


    def scatter(self):
        """ 元素打散 """
        # 先给每个数据赋一个随机数，再根据随机数排序
        # 也可以直接调用函数打散，但不一定很散
        random.shuffle(self.data)


    def store(self):
        """ 以CSV文件存储 """
        print('-------------- service is storing... total_rows={}'.format(len(self.data)), flush=True)
        rows = []
        now = get_cur_time()
        idx = 0
        for item in self.data:
            # 随机产生库存
            stock = generate_stock()
            row = [
                idx, item['cid'], item['cid2'],
                item['title'], item['author'], item['publisher'], item['publish_time'],
                item['price'], item['cur_price'], item['image'], item['info_url'], item['detail'],
                now, stock,
            ]
            rows.append(row)
            idx += 1

        headers = [
            'idx', 'cid', 'cid2', 'title', 'author', 'publisher', 'publish_time',
            'price', 'cur_price', 'images', 'info_url', 'detail', 'create_time',
            'stock'
        ]
        with open('data.csv','w', newline='', encoding='utf-8') as f:
            f_csv = csv.writer(f)
            f_csv.writerow(headers)
            f_csv.writerows(rows)


def run_service():
    service = Service(target_list=target_list)
    service.run()
    service.scatter()
    service.store()


def get_cur_time() -> str:
    localtime = time.localtime(time.time())
    return time.strftime('%Y-%m-%d %H:%M:%S', localtime)

def generate_stock() -> Number:
    return random.random(300, 2301)

if __name__ == "__main__":
    run_service()
