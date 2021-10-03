package CircleQuene

import "errors"

const QueneSize = 100

type CircleQuene struct {
	data  [QueneSize]interface{} //數據儲存結構
	front int                    //頭部位置
	rear  int                    //尾部位置
}

func Quenelength(q *CircleQuene) int { //隊列長度
	return (q.rear - q.front + QueneSize) % QueneSize

}

func InitQueue(q *CircleQuene) { //初始化,頭部尾部重合,為空
	q.front = 0
	q.rear = 0
}

func EnQueue(q *CircleQuene, data interface{}) (err error) {
	if (q.rear+1)%QueneSize == q.front%QueneSize {
		return errors.New("隊列已滿")
	}
	q.data[q.rear] = data             //入隊
	q.rear = (q.rear + 1) % QueneSize //循環
	return nil
}
func DeQueue(q *CircleQuene) (data interface{}, err error) {
	if q.rear == q.front {
		return nil, errors.New("隊列為空")
	}

	result := q.data[q.front]           //取出第一個數據
	q.data[q.front] = 0                 //清空數據
	q.front = (q.front + 1) % QueneSize //取餘數(循環)
	return result, nil
}
