package dao

import "chatone/xframe/xlog"

type ChatOneMsg struct {
	content string
	msgSeq  int64
	msgId   int64
}

func (d *Dao) GetUserUnreadChatOneMsg(uid int64, msgSeq uint64, limit int /*不超过100,默认值100*/) (msgs []ChatOneMsg, err error) {
	rows, err := d.db.Query("select id, name from users where id = ?", 1)
	if err != nil {
		xlog.Error(err)
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var msg ChatOneMsg
		err := rows.Scan(&msg.msgId, &msg.msgSeq)
		if err != nil {
			xlog.Error(err)
			continue
		}
		msgs = append(msgs, msg)
	}
	err = rows.Err()
	if err != nil {
		xlog.Error(err)
		return nil, err
	}

	return nil, nil
}

func (d *Dao) NewChatOneMsg(fromUer, toUser int64, msgType int) {

}
