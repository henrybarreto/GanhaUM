import React, {useEffect} from 'react';
import axios from 'axios'
import { useParams } from "react-router-dom";

import "./campaign.css"

export default function Campaign() {
    const { campaign_id } = useParams();

    useEffect(() => {
        axios.get("http://localhost:8080/api/public/campaigns/" + campaign_id).then((result) => {
            console.log(result.data)
        }).catch((error) => {
            console.error(error)
        })
    }, [])

    return (
    <div className="app">
      <header className="header">
        <div className="logo">
          <img className="logo-img" src="images/logo.svg" alt="GanhaUM's logo"/>
        </div>
      </header>

      <main className="main">
        <div className="campaign">
          <header className="campaign-header">
            <span className="campaign-label">Contribua e tenha a chance de receber:</span>
            <h1 className="campaign-title">XBOX LIVE GOLD 1 MÊS GLOBAL</h1>
          </header>

          <main className="campaign-main">
            <div className="campaign-thumbnail">
                <img className="campaign-thumbnail-img" src=""  alt="" />
            </div>

            <div className="campaign-info">
              <ul className="campaign-info-list">
                <li className="campaign-info-item">
                    <div className="campaign-info-item-icon">
                        <i className="ri-money-dollar-circle-fill"></i>
                    </div>
                    <span className="campaign-info-item-text">
                        <span className="campaign-info-item-label">Preço:</span>
                        <span className="campaign-info-item-value">R$ 1,00</span>
                    </span>
                </li>
                <li className="campaign-info-item">
                    <div className="campaign-info-item-icon">
                        <i className="ri-user-fill"></i>
                    </div>
                    <span className="campaign-info-item-text">
                        <span className="campaign-info-item-label">Contribuidores:</span>
                        <span className="campaign-info-item-value">34/40</span>
                    </span>
                </li>
                <li className="campaign-info-item">
                    <div className="campaign-info-item-icon">
                        <i className="ri-calendar-todo-fill"></i>
                    </div>
                    <span className="campaign-info-item-text">
                        <span className="campaign-info-item-label">Início:</span>
                        <span className="campaign-info-item-value">01/02/2023</span>
                    </span>
                </li>
                <li className="campaign-info-item">
                    <div className="campaign-info-item-icon">
                        <i className="ri-calendar-event-fill"></i>
                    </div>
                    <span className="campaign-info-item-text">
                        <span className="campaign-info-item-label">Fim:</span>
                        <span className="campaign-info-item-value">14/02/2023</span>
                    </span>
                </li>
              </ul>
            </div>

            <form id="campaign-form">
                <input id="campaign-id" type="hidden" name="campaign_id" value={campaign_id} />
            </form>
            <button className="campaign-button" type="submit" form="campaign-form" formMethod="get" formEncType="multipart/form-data" formAction="http://localhost:8080/api/public/contribute">Contribuir</button>
          </main>

            <footer className="campaign-footer">
                <div className="campaign-footer-description">
                    <h1 className="campaign-footer-description-title">Descrição</h1>
                    <p className="campaign-footer-description-text">
                        O Xbox Live Gold é um cartão de assinatura que é a sua chave para o mundo do entretenimento da Microsoft. Graças a este cartão poderá desfrutar de todos os privilégios da conta ouro. Isso inclui não apenas o acesso a uma das redes sociais mais robustas da indústria de jogos, mas também jogos com desconto, ofertas gratuitas exclusivas e a experiência multiplayer. Você não terá que se preocupar em perder o lançamento do seu título favorito, perder contato com amigos que fez durante uma partida intensa ou não ter dinheiro suficiente para comprar um clássico atemporal.
                    </p>
                </div>
            </footer>
        </div>
      </main>

      <footer className="footer">
        <div className="copyright">
          <p>Copyright © 2022</p>
        </div>
      </footer>
    </div>
  );
}
