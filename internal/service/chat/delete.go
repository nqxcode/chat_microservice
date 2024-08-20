package chat

import "context"

func (s *service) Delete(ctx context.Context, id int64) error {

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		errTx = s.messageRepository.DeleteByChatID(ctx, id)
		if errTx != nil {
			return errTx
		}

		errTx = s.chatToUserRepository.DeleteByChatID(ctx, id)
		if errTx != nil {
			return errTx
		}

		errTx = s.chatRepository.Delete(ctx, id)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
